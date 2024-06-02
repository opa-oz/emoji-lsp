# emoji-lsp

Emoji auto-completion!

[![GoDoc](https://pkg.go.dev/badge/github.com/opa-oz/emoji-lsp?status.svg)](https://pkg.go.dev/github.com/opa-oz/emoji-lsp?tab=doc)
[![Release](https://img.shields.io/github/release/opa-oz/emoji-lsp.svg?style=flat-square)](https://github.com/opa-oz/emoji-lsp/releases)

Say goodbye to emoji confusion and hello to seamless communication. 💻 
Emoji-LSP can be integrated seamlessly into your favorite text editor📄, providing Slack-like hints to effortlessly input emojis into your files. 
Whether you're adding a touch of personality to your code, spicing up your documents, or simply expressing yourself in a fun way! 
With intuitive suggestions built-in🔨 into your editor, never struggle to find the perfect emoji again. 

Boost your productivity and creativity today!

# Features

![markdownlinkcompletions](/docs/completion.png)


# Installation

## Neovim

1. Given Neovim access to the binary.
    - <details>
        <summary>Go install</summary>

        ```bash
            go install github.com/opa-oz/emoji-lsp@v0.0.4
        ```
    - [Mason.nvim](https://github.com/williamboman/mason.nvim) (from hosted binary)
    - Manually download from [releases](https://github.com/opa-oz/emoji-lsp/releases)

2. Modify your Neovim Configuration
```lua
vim.api.nvim_create_autocmd({ "BufReadPost" }, {
    pattern = { "*.md" }, -- specify target file extensions
    callback = function()
        vim.lsp.start({
            name = "emoji-lsp",
            cmd = { "emoji-lsp" },
            root_dir = vim.fn.getcwd(),
        })
    end,
})
```


----

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/S6S1UZ9P7)

