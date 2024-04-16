# Disable The Official League Of Legends Presence

## Summary 📜
League-RPC-Disabler is a Windows application designed to disable the official League of Legends Rich Presence.

## Description 📖
This application detects the launch of League of Legends. Upon detection, it modifies a specific JSON configuration file to disable the discord plugin.
This is something users want to do, especially when running software that enhances the rich presence from leage, such as: [league-rpc-linux](https://github.com/its-haze/league-rpc-linux)

<!---
Add image of league-rpc-linux here
-->

You could do all of this manually, but you would have to do it everytime you open league...
So having this in the background just saves a couple of minutes for you.

To do this manually, you would first need to find your plugin-manifest.json file.. usually it's in this location
```
C:\Riot Games\League of Legends\Plugins\plugin-manifest.json
```
and remove this entry:

```json
{
    "as": [],
    "name": "rcp-be-lol-discord-rp",
    "affinity": null,
    "lazy": false
}
```
Basically, all this program does, is to automate this process.


## Installation 🚀

1. Go to the [releases page](https://github.com/Its-Haze/league-rpc-disabler/releases/latest)
2. Download **league-rpc-disabler.exe**
3. Run the program


## Usage 🖱️

Just double click on the exe file, it's that simple!

The application will sit in the system tray and monitor for the League of Legends client process. When it detects the client, it will automatically modify the JSON configuration as specified.

---

## Local Development Installation 🛠️
To install League RPC Disabler, you need to have Go installed on your machine:

1. Install Go (if not installed): [https://golang.org/dl/](https://golang.org/dl/)
2. Clone the repository:
   ```
   git clone https://github.com/its-haze/league-rpc-disabler.git
   ```
3. Navigate into the project directory:
   ```
   cd league-rpc-disabler
   ```
4. Build the application:
   ```
   $env:GO111MODULE="on"; go build -ldflags "-H=windowsgui" -o "league-rpc-disabler.exe" ./cmd
   ```



<!---
Add image of tray icon here
-->

## Contributing 🤝

Contributions are what make the open-source community such a fantastic place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License ©️

Distributed under the MIT License. See `LICENSE` for more information.
