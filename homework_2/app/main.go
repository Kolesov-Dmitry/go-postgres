package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// fillDatabase fills database with "random" data using provided connection pool
func fillDatabase(ctx context.Context, pool *pgxpool.Pool) {
	// insert 50k records
	for i := 0; i < 50000; i++ {
		query := "insert into auth.users(first_name, last_name, email, pwd) values($1, $2, $3, $4)"

		_, err := pool.Exec(
			ctx,
			query,
			fmt.Sprintf("user %d", i+1),
			fmt.Sprintf("last_name %d", i+1),
			fmt.Sprintf("user_%d@mail.com", i+1),
			"123",
		)

		if err != nil {
			log.Printf("Error: unable to insert record: %s", err)
		}
	}
}

type LoadResults struct {
	Duration         time.Duration
	QueriesPerformed uint64
}

// loadTest performs database load test within a minute
func loadTest(ctx context.Context, dbpool *pgxpool.Pool, f func(ctx context.Context, dbpool *pgxpool.Pool) error) LoadResults {
	var queries uint64

	loadFunc := func(stopAt time.Time) {
		for {
			if err := f(ctx, dbpool); err != nil {
				log.Fatal(err)
			}

			atomic.AddUint64(&queries, 1)

			if time.Now().After(stopAt) {
				return
			}
		}
	}

	startAt := time.Now()
	stopAt := startAt.Add(1 * time.Minute)

	var wg sync.WaitGroup
	wg.Add(8)
	for i := 0; i < 8; i++ {
		go func() {
			loadFunc(stopAt)
			wg.Done()
		}()
	}
	wg.Wait()

	return LoadResults{
		Duration:         time.Now().Sub(startAt),
		QueriesPerformed: queries,
	}
}

// selectUserByEmail using for "select user by email" query load test
func selectUserByEmail(ctx context.Context, dbpool *pgxpool.Pool) error {
	query := "SELECT id, first_name, last_name, email, pwd, created_at FROM auth.users WHERE email = $1"
	rows, err := dbpool.Query(ctx, query, "user_100@email.com")
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

func main() {
	// databsae connection string
	url := "postgres://blog_srv:asdfg@localhost:5432/blog"

	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Fatal(err)
	}

	// configure connection pool
	cfg.MaxConns = 2
	cfg.MinConns = 1

	cfg.HealthCheckPeriod = 1 * time.Minute
	cfg.MaxConnLifetime = 24 * time.Hour
	cfg.MaxConnIdleTime = 30 * time.Minute
	cfg.ConnConfig.ConnectTimeout = 1 * time.Second

	// connect to database
	pool, err := pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	// fillDatabase(context.Background(), pool)

	// start tests
	res := loadTest(context.Background(), pool, selectUserByEmail)

	fmt.Println("duration:", res.Duration)
	fmt.Println("queries:", res.QueriesPerformed)

	qps := res.QueriesPerformed / uint64(res.Duration.Seconds())
	fmt.Println("QPS:", qps)
}
