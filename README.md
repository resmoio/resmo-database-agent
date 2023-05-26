# Resmo Database Agent
[Resmo](https://www.resmo.com/) Database agent is an open database agent for your on premise database assets.
After you create and integration and setting the configurations as specified in the How it works section, database agent starts to operate.
Once triggered, it sends requests to our servers, pulling resources inside the database. Then, we process them on the server.
You can use the database agent as a binary or docker image depending on the installation path you choose.
Resmo performs the resource validation with a unique IngestKey. Note: An IngestKey is different for each integration.

## Supported Databases:
- [x] PostgreSQL
- [x] MySQL
- [x] ClickHouse
- [x] MongoDB

## Documentation

For full documentation, visit [docs.resmo.com](https://docs.resmo.com/product/getting-started/readme)

## How it works

1. Login to your Resmo account, and go to Integrations page. 
2. Click on the desired database integration.
3. Type a descriptive name for the integration and optionally a description.
4. Click on create.
5. Go to [Resmo-Database-Agent Releases](https://github.com/resmoio/resmo-database-agent/releases).
6. If you want to binary version, click on the desired OS version of the Resmo Database Agent under the Assets.
7. After downloading you can run the database agent via providing the config values.
8. Available configuration flags are:
    * ingestKey: The secret created for your integration at Resmo (**Required**). 
    * dsn: The datasource URL (**Required**). We are expecting the URL to follow a specific format:
      * [MySQL](https://github.com/go-sql-driver/mysql)
        * ```mysql://user:password@tcp(host:port)/database```
        * ```mysql://user:password@unix(/path/to/socket)/database```
      * [PostgreSQL](https://github.com/lib/pq)
          * ```postgres://user:password@host:port/database```
          * ```postgresql://user:password@host:port/database```
      * [ClickHouse](https://github.com/ClickHouse/clickhouse-go)
          * ```clickhouse://user:password@host:port?database=database```
      * [MongoDB](https://github.com/mongodb/mongo-go-driver)
          * ```mongodb://user:password@host:port/?authSource=database``` 
    * schedule: Schedule for running queries, "10m" for 10 minutes schedule for example.
    * timeout: Timeout duration for database connections, ingesting etc. 
    * dbIdentifier: Database identifier for related Resmo resources.
9. You can run database runner for the binary version with commands:
   ```bash
    $ ./resmo-database-agent -ingestKey=<RESMO_INGEST_KEY> -dsn="clickhouse://username:password@localhost:8123"
   ```
   You can also supply the secret values as environment variables in the following form: :
   ```bash
   $ RESMO_INGEST_KEY=<RESMO_INGEST_KEY> 
   $ DSN="clickhouse://username:password@localhost:8123" 
   $ ./resmo-database-agent
   ```
10. You can also run the agent as a Docker container as follows:
    ```bash
    $ docker run ghcr.io/resmoio/resmo-database-agent -ingestKey=<RESMO_INGEST_KEY> -dsn="clickhouse://username:password@localhost:8123"
    ```
     or as environment variables:
    ```bash 
    $ docker run -e DSN="clickhouse://username:password@localhost:8123" -e RESMO_INGEST_KEY=<RESMO_INGEST_KEY> /resmo-database-agent
    ```
11. You can make use of Kubernetes CronJob to schedule the job instead of the keeping the agent open all the time. 
Following is an example kubernetes manifest that will connect to your database and Kubernetes scheduler ensures container runs in every 10 minutes:
```yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: my-cronjob
spec:
  schedule: "*/10 * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
            - name: my-cronjob
              image: ghcr.io/resmoio/resmo-database-agent:<LATEST_VERSION>
              env:
                - name: DSN
                  value: <DSN>
                - name: RESMO_INGEST_KEY
                  value: <RESMO_INGEST_KEY>
          restartPolicy: OnFailure
```    
12. You are ready! Now you can start querying your database resources!

## Support

For any questions, you can reach us from [here.](https://www.resmo.com/contact)
