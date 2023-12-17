## Генераторы кода на Go

Недавно изучал генерацию кода на языке Go и написал несколько собственных генераторов. Вот некоторые из них:

### Генератор последовательности чисел Фибоначчи

Этот генератор создает простую программу для вычисления чисел Фибоначчи с использованием рекурсии.

### Создатель пустых структур

Генератор, который создает пустые структуры.

### Генератор программы "Hello, World!"

Этот генератор создает программу, выводящую на экран приветствие "Hello, World!".

## Использование

Для генерации программ выполните следующие действия:

1. Для генератора кода Фибоначчи:
    ```sh
    go run fibonacci_generator.go
    ```

2. Для создания пустой структуры:
    ```sh
    go run empty_struct_creator.go
    ```

3. Для генератора программы "Hello, World!", либо:
    - Выполните:
        ```sh
        go run hello_world_generator.go
        ```
    - Или добавьте следующий комментарий в начало файла:
        ```go
        //go:generate go run hello_world_generator.go
        ```
      Затем перейдите в расположение файла и выполните:
      ```sh
      go generate
      ```


