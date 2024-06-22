package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/grafana/pyroscope-go"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/younesious/go-microservices/authentication/data"
)

const (
	webPort         = "8083"
	metricsPort     = "9090"
	jaegerAgentPort = "6831"
	pyroscopePort   = "4040"
)

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("---------------------------------------------")
	log.Println("Attempting to connect to Postgres...")

	conn := connectToDB()
	if conn == nil {
		log.Panic("can't connect to postgres!")
	}

	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	go startMetricsServer()
	go startJaegerTracer()
	go startPyroscopeProfiler()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	log.Printf("Starting authentication service on port %s\n", webPort)
	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}

	select {} // Prevent the main function from exiting
}

func startMetricsServer() {
	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Starting metrics server on port %s\n", metricsPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", metricsPort), nil))
}

func startJaegerTracer() opentracing.Tracer {
	tracer, closer := initJaeger("auth-service")
	defer closer.Close()
	return tracer
}

func startPyroscopeProfiler() {
	pyroscope.Start(pyroscope.Config{
		ApplicationName: "auth-service",
		ServerAddress:   fmt.Sprintf("http://pyroscope:%s", pyroscopePort),
		Logger:          pyroscope.StandardLogger,
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})
	select {} // Prevent the goroutine from exiting
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	count := 0
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not ready...")
			count++
		} else {
			log.Println("Connected to database!")
			return connection
		}

		if count > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg := config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: fmt.Sprintf("jaeger:%s", jaegerAgentPort),
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		log.Fatalf("cannot initialize Jaeger Tracer: %v", err)
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}
