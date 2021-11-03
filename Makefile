.SILENT:

mongo:
	docker run --name stock_picker -v .:/data/db -d mongo