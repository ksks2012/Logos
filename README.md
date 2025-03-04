# Database code generation

## go-literal-code-gen

- install
    
    ```
    go install github.com/yinyin/go-literal-code-gen@latest
    ```
    
- genertate
    
    ```
    go generate ./internal/dao/dbversion/mysql
    ```

# Command

## Testing

- Run tests

    ```
    go test ./testing/...
    ```

# TODO

- [] Show info
- [] Save & Load
- [] Generator
    - [] Character
    - [] Global map
- [] Global update
- [] Timer
    - [] Run
    - [] Pause
    - [] Time multiplier

## info

- [x] List of units
- [x] Simple logging system
- [] Simple ASCII visualization of maps
- [] Simple NPC behavior change tracking
- [] JSON / YAML to store game state
    - [] JSON
        - [x] Save
        - [] Load
    - [] YAML
        - [] Save
        - [] Load