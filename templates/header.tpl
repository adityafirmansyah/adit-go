{{define "header"}}<!DOCTYPE html>
<html id="top">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="chrome=1">

    <title>{{ .Title }}</title>

    <link rel="stylesheet" href="http://fonts.googleapis.com/css?family=Gentium+Book+Basic:300,400,700|Lato:300,400">
    <link rel="stylesheet" href="/css/main.css">

</head>

<body>

    <header class="Header Bread">
        <a href="/"><img class="Header-avatar" src="/img/avatar.jpg" alt="adit's avatar"></a>
        <a class="Header-hamburger" href="#top"></a>
        <div class="Header-target">
            <h1 class="Header-title">Aditya Firmansyah</h1>
            <p class="Header-description">Just a simple Blog</p>

            <nav class="Nav">
                <div class="Nav-group">
                    <a class="Nav-item " href="/#">Projects</a>
                    <a class="Nav-item " href="/#">Lab</a>
                    <a class="Nav-item is-active" href="/#">Blog</a>
                </div>
                <div class="Nav-group">
                    <a class="Nav-item " href="/about/">About</a>
                </div>
            </nav>

        </div>
    </header>

    <main class="Post">
{{end}}
