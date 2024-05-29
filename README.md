# THAG Stack Webapp Template (TailwindCSS + HTMX + Alpine.js + Go)
A Web application template to show off the THAG Stack.

## Development Setup

For the best development experience it is recomended to use a `Linux` or `MacOS` operating system and [VSCode](https://code.visualstudio.com) or [VSCodium - Open version of VSCode](https://vscodium.com/) as the IDE.

### 1. Install Golang v1.22 or greater

Follow the instalation instructions [here](https://go.dev/doc/install).

### 2. Install [Air](https://github.com/cosmtrek/air)

Air makes it easy do develop Go applications by automatically recompiling on code changes.

To install it do the following:

1. Create a directory named `air` within the `dev` directory:
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

1. Create a directory named `templ` within the `dev` directory:
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

1. Create a directory named `tailwindcss` within the `dev` directory:
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

### 5. Setup IDE for development

While it's not strictly necessary, adding intellisense support for `HTMX`, `Templ`, `TailwindCSS` and `Alpine.js` in your prefered IDE (if avaliable), will result in a much richer development experience.

Here is a list of extencions for VSCode/VSCodium recommended for this project:
- [Go](https://marketplace.visualstudio.com/items?itemName=golang.go)
- [templ-vscode](https://marketplace.visualstudio.com/items?itemName=a-h.templ)
- [HTMX Attributes](https://marketplace.visualstudio.com/items?itemName=CraigRBroughton.htmx-attributes)
- [Tailwind CSS IntelliSense](https://marketplace.visualstudio.com/items?itemName=bradlc.vscode-tailwindcss)
- [Alpine.js IntelliSense](https://marketplace.visualstudio.com/items?itemName=adrianwilczynski.alpine-js-intellisense)


On VSCodium some extensions are not avaliable by default on the extensions tab (`Alpine.js IntelliSense` for examle), but they can still be installed manually. 

To manually install an extension on VSCodium, search for it on [VSCode Marketplace](https://marketplace.visualstudio.com) and look the download link for the latest version. After the `vsix` extension file finishes downloading, install it by running:
```
codium --install-extension extension_file_name.vsix

# replace 'extension_file_name' for the actual file name of the extension
```

## Development mode

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