# Folder11-auto

## ⚠️⚠️⚠️ Please note: This project is incomplete and may cause irreversible consequences on files. Please do not run it casually ⚠️⚠️⚠️

`Folder11-auto` is an automation script written in Go, designed to automatically set custom icons for Windows folders. The tool uses `desktop.ini` files to specify the icon for each folder and can recursively traverse folders to automatically assign the best matching icon.

## Features

- Automatically assigns icons to folders, with recursive support for subfolder icons.
- Supports icon resources from a GitHub submodule or website scraping.
- Customizable icon directory and recursion depth.
- Uses fuzzy matching between folder names and icon names to assign the most suitable icon.
- Supports custom tokenization rules and keyword weight configuration files.

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/your-repo/Folder11-auto.git
cd Folder11-auto
git submodule update --init --recursive
```

### 2. Install dependencies

Make sure you have the Go environment installed, then run the following command to install the project dependencies:

```bash
go mod tidy
```

### 3. Build the project

```bash
go build -o Folder11-auto.exe
```

## Usage

### 1. Run the script

Run `Folder11-auto.exe` to set custom icons for the specified folder. You can use the following command-line arguments:

```
Usage: Folder11-auto.exe [options] [folderPath]
Options:
  -f, --folder        Specify the target folder path
  -r, --recursive     Set the recursion depth (default: 3)
  -s, --source        Specify the icon source ("Folder11-Ico" or "website")
  -i, --icon-dir      Specify the icon directory (default: "./Folder11-Ico/ico")
  -h, --help          Display help information
```

#### Examples:

1.  Set a custom icon for the target folder `E:\your_folder` with a recursion depth of 4, and icon source as `Folder11-Ico`:

```bash
Folder11-auto.exe -f E:\your_folder -r 4 -s Folder11-Ico
```

1.  Set an icon for the folder by scraping icons from a website:

```bash
Folder11-auto.exe -f E:\your_folder -r 4 -s website
```

### 2. Argument explanation

-   `-f, --folder`: Specifies the target folder path (required).

-   `-r, --recursive`: Sets the recursion depth, with the default value being 3. It determines the maximum depth for recursive subfolder processing.

-   ```
    -s, --source
    ```

    : The source of icons, with two possible values:

    -   `Folder11-Ico`: Get icons from the project submodule's icon folder.
    -   `website`: Scrape icons from a specified website (this feature requires a web scraper).

-   `-i, --icon-dir`: Specifies the icon directory path, default is `./Folder11-Ico/ico`.

### 3. Icon matching

The script will use folder names (tokenized by spaces, underscores, hyphens, etc.) and match them with icon file names. The system will select the best matching icon, prioritizing those with longer names and smaller edit distances.

### 4. `desktop.ini` Configuration

The script modifies the `desktop.ini` file to set the folder icon. Make sure that the target folder contains a `desktop.ini` file with the correct icon file path. For example:

```ini
[.ShellClassInfo]
IconFile=C:\path\to\icon.ico
IconIndex=0
Attributes=2
```

### 5. Set `desktop.ini` as a system file

For the folder icon to take effect, make sure that the `desktop.ini` file is marked as a system and hidden file. You can do this using the following command:

```bash
attrib +s +h "E:\your_folder\desktop.ini"
```

## Contributing

Feel free to raise issues and submit pull requests!

## License

This project is licensed under the MIT License. See the [LICENSE](https://chatgpt.com/c/LICENSE) file for details.



## ⚠️⚠️⚠️请注意，当前项目并未完成且可能对文件造成不可恢复的后果，请勿随意运行⚠️⚠️⚠️

`Folder11-auto` 是一个用 Go 编写的自动化脚本，旨在为 Windows 文件夹自动设置自定义图标。该工具使用 `desktop.ini` 文件来指定每个文件夹的图标，可以支持递归地遍历文件夹并自动为文件夹分配最匹配的图标。

## 功能特性

- 自动为文件夹分配图标，支持递归设置子文件夹图标。
- 支持图标资源来自 GitHub 子模块或网站爬取。
- 支持自定义图标目录和递归深度。
- 根据文件夹名称与图标名称进行模糊匹配，为文件夹设置最匹配的图标。
- 配置文件支持自定义分词规则和关键词权重。

## 安装

### 1. 克隆项目

```bash
git clone https://github.com/your-repo/Folder11-auto.git
cd Folder11-auto
git submodule update --init --recursive
```

### 2. 安装依赖

确保你已经安装了 Go 环境，运行以下命令安装项目依赖：

```bash
go mod tidy
```

### 3. 构建项目

```bash
go build -o Folder11-auto.exe
```

## 使用说明

### 1. 运行脚本

运行 `Folder11-auto.exe` 来为指定文件夹设置自定义图标。可以使用以下命令行参数：

```
Usage: Folder11-auto.exe [options] [folderPath]
Options:
  -f, --folder        指定目标文件夹路径
  -r, --recursive     指定递归深度 (默认: 3)
  -s, --source        指定图标来源 ("Folder11-Ico" 或 "website")
  -i, --icon-dir      指定图标文件夹路径 (默认: "./Folder11-Ico/ico")
  -h, --help          显示帮助信息
```

#### 示例：

1.  为目标文件夹 `E:\your_folder` 设置自定义图标，递归深度为 4，图标来源为 `Folder11-Ico`：

```bash
Folder11-auto.exe -f E:\your_folder -r 4 -s Folder11-Ico
```

1.  使用爬取网站图标的方式为文件夹设置图标：

```bash
Folder11-auto.exe -f E:\your_folder -r 4 -s website
```

### 2. 参数说明

-   `-f, --folder`：指定目标文件夹路径，必须提供。

-   `-r, --recursive`：递归深度，默认值为 3。指定递归子文件夹的最大深度。

-   ```
    -s, --source
    ```

    ：图标来源，支持两个值：

    -   `Folder11-Ico`：从项目子模块的图标文件夹中获取图标。
    -   `website`：从指定网站爬取图标（该功能依赖网络爬虫）。

-   `-i, --icon-dir`：指定图标文件夹路径，默认值为 `./Folder11-Ico/ico`。

### 3. 图标匹配

脚本会根据文件夹名称（通过分词如空格、下划线、连字符等）与图标文件名称进行模糊匹配。系统会选择最佳匹配的图标，优先选择名称长度较长、编辑距离较小的图标。

### 4. `desktop.ini` 配置

该脚本通过编辑 `desktop.ini` 文件来为文件夹设置图标。确保目标文件夹中存在 `desktop.ini` 文件，并且其内容包括正确的图标文件路径。例如：

```ini
[.ShellClassInfo]
IconFile=C:\path\to\icon.ico
IconIndex=0
Attributes=2
```

### 5. 设置 `desktop.ini` 文件为系统文件

为了使文件夹图标生效，确保 `desktop.ini` 文件被标记为系统文件和隐藏文件。可以使用以下命令：

```bash
attrib +s +h "E:\your_folder\desktop.ini"
```

## 贡献

欢迎提出问题和提交 Pull Request！

## 许可证

本项目采用 MIT 许可证，详见  [LICENSE](LICENSE)  文件。

