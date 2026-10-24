@app/AGENTS.md

# Web Journal

This is a project to build a website where users tell stories or share experiences by placing pictures, videos, text, and other widgets onto a page with a visually appealling layout.

## General Guidelines

This project is largely a learning and practice opportunity for me. I want to experiment with some backend web service architecture ideas for the API portion and get some more hands on experience with building a react based web site. Therefore, only make changes yourself if I explicitly tell you. Keep the scope of your work very focused. If you have questions, ask me. Don't make assumptions. 

The exception is in the `api/cmds` directory. I will have you build maintenance and management commands for me using the API's services.

## Maintenance Commands

Maintenance commands are placed in the `api/cmds` directory. They use services to interact with the web service data. They are interactive CLI tools. Instead of taking a bunch of arguments, prompt the user for each piece of information required and validate each one. Commands provide rich output using colors and tables and other structures to be as friendly as possible. They always ask for confirmation before taking potentially destructive actions.
