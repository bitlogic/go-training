# Installation and Setup

![setup](/go-training-beginner/getting_starting/img/beginners.png)

## Descarga

Para empezar descargamos Go desde la sección de [descarga](https://go.dev/dl/) de la web oficial de Go.

![download](/go-training-beginner/getting_starting/img/download.png)

## Instalación

Las siguientes instrucciones son de la [web oficial](https://go.dev/doc/install) de Go

### MacOS

1. Open the package file you downloaded and follow the      prompts to install Go. The package installs the Go distribution to `/usr/local/go`. The package should put the `/usr/local/go/bin` directory in your `PATH` environment variable. You may need to restart any open Terminal sessions for the change to take effect.

2. Verify that you've installed Go by opening a command prompt and typing the following command:

```cmd
$ go version
```

3. Confirm that the command prints the installed version of Go.

### Linux

1. Remove any previous Go installation by deleting the `/usr/local/go` folder (if it exists), then extract the archive you just downloaded into `/usr/local`, creating a fresh Go tree in `/usr/local/go`:

```cmd
$ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.1.linux-amd64.tar.gz
```

>You may need to run the command as root or through sudo)

**Do not untar** the archive into an existing `/usr/local/go `tree. This is known to produce broken Go installations.

2. Add `/usr/local/go/bin` to the `PATH` environment variable. You can do this by adding the following line to your `$HOME/.profile` or `/etc/profile` (for a system-wide installation):

```cmd
export PATH=$PATH:/usr/local/go/bin
```

**Note**: Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source `$HOME/.profile`.

3. Verify that you've installed Go by opening a command prompt and typing the following command:

```cmd
$ go version
```

4. Confirm that the command prints the installed version of Go.

### Windows

1. Open the MSI file you downloaded and follow the prompts to install Go.

By default, the installer will install Go to Program Files or Program Files (x86). You can change the location as needed. After installing, you will need to close and reopen any open command prompts so that changes to the environment made by the installer are reflected at the command prompt.

2. Verify that you've installed Go.
    1. In Windows, click the Start menu.
    2. In the menu's search box, type cmd, then press the Enter key.
    3. In the Command Prompt window that appears, type the following command:

```cmd
$ go version
```

3. Confirm that the command prints the installed version of Go.

### VS Code

Go tiene su propio IDE (golang) que es muy bueno, pero por ahora recomiendo usar VS Code ya que lo mas probable es que ya estés familiarizado con este, pero lo dejo a criterio de cada uno.

Podemos descargar VS Code desde este [link](https://code.visualstudio.com/download)

![VS Code](/go-training-beginner/getting_starting/img/vscode.png)

### Extension

Es muy útil tener instalado la [extensión de go](https://code.visualstudio.com/docs/languages/go) para VS Code

![extensión](/go-training-beginner/getting_starting/img/extension.png)