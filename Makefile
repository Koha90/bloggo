build:
	@go build -o bin/bloggo cmd/main.go

run: build
	@./bin/bloggo &
	@npm install --prefix ./frontend
	@npm run dev --prefix ./frontend

migrateup:
	@sqlite3 database/blog.db < migrations/000001_create_users_table.up.sql

migratedown:
	@sqlite3 database/blog.db < migrations/000001_create_users_table.down.sql
