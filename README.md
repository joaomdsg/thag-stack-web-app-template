# THAG Stack Webapp Template (TailwindCSS + HTMX + Alpine.js + Go)
A Web application template to show off the THAG Stack.

## Setup

### 1. Install Golang v1.22 or greater

Follow the instalation instructions [here](https://go.dev/doc/install).

### 2. Install [Air](https://github.com/cosmtrek/air)

Air makes it easy do develop Go applications by automatically recompiling on code changes.

To install it do the following:

1. Create a directory named `air` that should be inside another directory named `dev`:
    ```
    mkdir -p ./dev/air
    ```

2. Download air standalone executable from [here](https://github.com/cosmtrek/air/releases).

3. Move the executable to  `./dev/air`.

4. Rename the executable to `bin`.

5. Give it executable permissions:
    ```
    chmod +x ./dev/air/bin
    ```

### 3. Install [Templ](https://templ.guide/)

Templ is a powerful HTML templating tool that makes web development with Go easier.

To install it do the following:

To install it do the following:

1. Create a directory named `templ` that should be inside another directory named `dev`:
    ```
    mkdir -p ./dev/templ
    ```

2. Download templ standalone executable tar ball from [here](https://github.com/a-h/templ/releases).

3. Uncompress it:
    ```
    tar xfz path/to/templ.tar.gz
    ```

4. Move the executable to  `./dev/templ`.

5. Rename the executable to `bin`.

6. Give it executable permissions:
    ```
    chmod +x ./dev/templ/bin
    ```

### 4. Install [TailwindCSS](https://tailwindcss.com/)

TailwindCSS allows for modern and beautiful web UIs without leaving the HTML.

To install it do the following:

1. Create a directory named `tailwindcss` that should be inside another directory named `dev`:
    ```
    mkdir -p ./dev/tailwindcss
    ```

2. Download tailwind standalone executable from [here](https://github.com/tailwindlabs/tailwindcss/releases/).

3. Move the executable to  `./dev/tailwindcss`.

4. Rename the executable to `bin`.

5. Give it executable permissions:
    ```
    chmod +x ./dev/tailwindcss/bin
    ```

## Development

Start the app in development mode by running:
```
make dev
```
Open a web browser and navigate to the URL http://localhost:3333. The application will automatically update as you make changes to the codebase.


## Build for production
Build a production ready executable by running:
```
make build
```
The build executable will be avaliable on `./tmp/main`.