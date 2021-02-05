---
id: aws
title: Amazon Web Services Cli
sidebar_label: AWS
---

## What

Displays the current active AWS Cli Profile and Region

## Sample Configuration

```json
{
    "type": "aws",
    "style": "powerline",
    "powerline_symbol": "",
    "invert_powerline": false,
    "foreground": "#100e23",
    "background": "#ffffff",
    "leading_diamond": "",
    "trailing_diamond": "",
    "properties": null
},
```

## Properties

- display_profile_name: `boolean` - hides or shows the profile name - defaults to `true`
- display_region: `boolean` - hides or shows the region - defaults to `true`
- separator: `string` - separator text/icon displayed between the profile name and region - defaults to an organe amazon logo: `<#F8991D>\uf52c</>`
