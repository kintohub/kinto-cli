<p align="center">
  <a>
    <img src="https://www.kintohub.com/logo-full-color-horizontal.svg" href="https://kintohub.com" height="40"/>
  </a>
  <h2 align="center">Kinto CLI</h2>
  <p align="center">The official Commandline interface for KintoHub</p>
  <p align="center">
    <img src="https://img.shields.io/badge/status-Alpha-green?style=for-the-badge"></img>
    <img src="https://img.shields.io/github/v/release/kintohub/kinto-cli?style=for-the-badge"></img>
  </p>
</p>

# :sparkles: Features

- **Easy Access**

  - One click access to all your environments and services.
  - Check status of all your deployed services on KintoHub.

- **Dev tools**
  - Port forward your remote services to your local machine for real-time debugging and preview using the `teleport` feature.


# :computer: Installation


## Linux / MacOS

### :seedling: Requirements :

- Make sure you use bash to run this script.
- `unzip` should be installed. If you don't already have it install it using `sudo apt-get install -y unzip` or similar command depending on your linux distro.

Run the installation script to install kinto CLI to its default location `usr/local/bin`

```
curl -L https://cli.kinto.io/install.sh | bash
```

## Windows

### :seedling: Requirements :

- Windows 7 SP1+ / Windows Server 2008+
- [PowerShell 5](https://aka.ms/wmf5download) (or later, include [PowerShell Core](https://docs.microsoft.com/en-us/powershell/scripting/install/installing-powershell-core-on-windows?view=powershell-6))
- PowerShell must be enabled for your user account e.g. `Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser`
- Run PowerShell as Administrator to avoid any errors.

Run the following command from your PowerShell to install Kinto CLI to its default location `C:\Users\<user>\kinto`

```
Invoke-Expression (New-Object System.Net.WebClient).DownloadString('https://cli.kinto.io/install.ps1')
```

OR

```
iwr -useb cli.kinto.io/install.ps1 | iex
```
---

You can also download the latest available release for your Operating System from [releases](https://github.com/kintohub/kinto-cli/releases) and add it to your global `PATH` manually.

---

## :mag: Commands

For a detailed overview of the available CLI commands see [docs](https://docs.kintohub.com/anatomy/cli#commands).

## :alembic: Contributing

See the [Contributing Guide](./CONTRIBUTING.md).

## :books: Documentation

Check out the official KintoHub [Documentation](https://docs.kintohub.com/)

## :ambulance: Support

**Contact us:** https://www.kintohub.com/contact-us

**Discord:** https://discordapp.com/invite/E2CMjKP
