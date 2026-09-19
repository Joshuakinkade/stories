@app/AGENTS.md

# Web Journal

This is a project to build a website where users tell stories or share experiences by placing pictures, videos, text, and other widgets onto a page with a visually appealling layout.

## General Guidelnes

This project is largely a learning and practice opportunity for me. I want to experiment with some backend web service architecutre ideas for the api portion and get some more hands on experience with building a react based web site. Therefore, do *not* edit files. If I ask questions about how to do something, give me ideas and point me in the right direction, but don't do the work for me.

The exception is in the `api/cmds` directory. I will have you build maintenance and management commands for me using the services that I provide.

## Maintenance Commands

Maintenance commands are placed in the `api/cmds` directory. They use services to interact with the web service data. They are interactive cli tools. Instead of taking a bunch of arguments, prompt the user for each peice of information required and validate each one. Commands provide rich output using colors and tables and other structures to be as friendly as possible. They always aks for confirmation before taking potentially destructive actions.
