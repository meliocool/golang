# XSS (Cross Site Scripting)

## Introduction
- XSS is one of the many security issue faced when making web
- XSS is when someone injects JavaScript in a Parameter that will be rendered in the web browser
- In most cases the goal of XSS is to get Cookies of users visiting the web, to take over their account

## Auto Escape
- Go Template has a feature called Auto Escape
- It can detect data that should be displayed in the template
- If it has tag tag html or JavaScript code, it will automatically escape