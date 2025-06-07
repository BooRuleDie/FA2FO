# Connection & Security

Redis provides several commands for managing connections and implementing security measures:

## Basic Connection Commands
- `PING`: Tests connectivity. Server responds with "PONG" if alive
- `ECHO message`: Returns the specified message back to client
- `SELECT index`: Switch to specified database (0-15 by default)

## Durability & Persistence
Redis offers multiple persistence options:
- RDB (Redis Database): Point-in-time snapshots
  ```bash
  # Configure RDB snapshots every 60 seconds if at least 100 changes
  CONFIG SET save "60 100"

  # Force RDB snapshot
  SAVE

  # Check last save status
  LASTSAVE
  > (integer) 1639516800
  ```

- AOF (Append Only File): Logs every write operation
  ```bash
  # Enable AOF
  CONFIG SET appendonly yes

  # Set AOF fsync policy
  CONFIG SET appendfsync everysec

  # Check AOF status
  INFO persistence
  > # Persistence
  > aof_enabled:1
  > aof_rewrite_in_progress:0
  > aof_last_write_status:ok
  ```

- RDB+AOF: Combined approach for enhanced durability
  ```bash
  # Enable both RDB and AOF
  CONFIG SET save "60 100"
  CONFIG SET appendonly yes

  # Verify both are enabled
  INFO persistence
  > # Persistence
  > rdb_last_save_time:1639516800
  > aof_enabled:1
  ```

## Client Management
- `CLIENT LIST`: Shows information about connected clients
- `CLIENT KILL id client-id`: Terminates connection of specified client

## Authentication & Security
- `CONFIG SET requirepass "password"`: Sets authentication password
- `AUTH password`: Authenticates client using specified password

Best practices:
- Always use strong passwords
- Enable protected mode in production
- Configure proper persistence strategy
- Monitor client connections
