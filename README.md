# Fiber framework

### Makefile

1. Add GOROOT/Bin to PATH
   ```sh
   $ echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.zshrc
   $ source ~/.zshrc
   ```
2. Run Make commands
   ```sh
   $ make watch
   ```
3. Run PostgreSQL instance via docker container
   ```sh
   $ docker-compose up -d
   ```
4. Test access database postgres
   ```sh
   $ psql -U <USER_NAME> -p <PORT> -h <HOST> -d <DB_NAME>
   ```
5. 
