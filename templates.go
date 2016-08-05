package main

import "html/template"

var powerT = template.Must(template.New("power").Parse(`{
  "text": "{{.title}}",
  "attachments": [
      { "text": "{{.status}}" },
      {
          "fallback": "Your slack client does not support interactivity :(",
          "callback_id": "{{.callback}}",
          "color": "#0f0",
          "actions": [
            { "name": "power", "text": "Status", "type": "button", "value": "status" },
            { "name": "power", "text": "On", "type": "button", "value": "on", "style": "primary" }
          ]
      },
      {
          "fallback": "Your slack client does not support interactivity :(",
          "callback_id": "{{.callback}}",
          "color": "#f00",
          "actions": [
            { "name": "power", "text": "Off", "type": "button", "value": "off", "style": "danger" },
            { "name": "power", "text": "Reset", "type": "button", "value": "reset", "style": "danger" },
            { "name": "power", "text": "Cycle", "type": "button", "value": "cycle", "style": "danger" }
        ]
    }
  ]
}`))
