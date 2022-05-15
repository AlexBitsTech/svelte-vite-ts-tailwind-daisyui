# README

## About

This is a wails template based on the official Wails Svelte-TS template.

It implements Svelte, Vite, TypeScript, Tailwindcss and the FlowBite framework.

## Installation

To use the template, run `wails init` with the url of this repo as the template parameter like so:
```
wails init -n <name_of_project> -t https://github.com/AlexBitsTech/svelte-vite-ts-tailwind-flowbite
```

## Updating dependencies

If you wish to update the packages in this module you can either update each one manually using `npm update [-g] [<pkg>...]` or use the `npm-check_update` package to update all packages at once.

To use the latter, you must install it globally by running :
```
npm i -g npm-check-update
```
 You can then check which packages need an update by running:
 ```
 ncu
 ```
 and update all the modules using:
 ```
 ncu -u
 ```
 This command will only update the versions in the `package.json` file. You still need to run:
 ```
 npm install
 ```
 to do the actual upgrade.

## Live Development

To run in live development mode, run `wails dev` in the generated project directory.

## Building

To build a redistributable, production mode package, use `wails build`.

## Disclaimer

This template has been tested with the following versions of Wails:
* `v2.0.0-beta.36`

## Licence