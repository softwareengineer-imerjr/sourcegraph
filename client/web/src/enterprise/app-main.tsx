// This is the entry point for the enterprise Sourcegraph app

// Order is important here
// Don't remove the empty lines between these imports

// prettier-ignore-start
import '@sourcegraph/shared/src/polyfills'
// prettier-ignore-end

import '../initBuildInfo'
import '../monitoring/initMonitoring'

import { createRoot } from 'react-dom/client'

import { logger } from '@sourcegraph/common'

import { initAppShell } from '../storm/app-shell-init'

import { EnterpriseWebApp } from './EnterpriseWebApp'

import { Command } from '@tauri-apps/api/shell'

// Start Backend early
const command = Command.sidecar("../.bin/backend", [], {})
command.on('close', data => {
    console.log(`backend command finished with code ${data.code} and signal ${data.signal}`)
})
command.on('error', error => console.log(`command error: "${error}"`))
command.stdout.on('data', line => console.log(`backend out: ${line}`))
command.stderr.on('data', line => console.log(`backend err: ${line}`))
const child = await command.spawn()

window.addEventListener('backend-message', async (event) => {
    console.log(`msg: ${event}`)
})

console.log(`backend started with pid ${child.pid} `)

const appShellPromise = initAppShell()

// It's important to have a root component in a separate file to create a react-refresh boundary and avoid page reload.
// https://github.com/pmmmwh/react-refresh-webpack-plugin/blob/main/docs/TROUBLESHOOTING.md#edits-always-lead-to-full-reload
window.addEventListener('DOMContentLoaded', async () => {
    const root = createRoot(document.querySelector('#root')!)

    try {
        const { graphqlClient, temporarySettingsStorage } = await appShellPromise

        root.render(
            <EnterpriseWebApp graphqlClient={graphqlClient} temporarySettingsStorage={temporarySettingsStorage} />
        )
    } catch (error) {
        logger.error('Failed to initialize the app shell', error)
    }
})

if (process.env.DEV_WEB_BUILDER === 'esbuild' && process.env.NODE_ENV === 'development') {
    new EventSource('/.assets/esbuild').addEventListener('change', () => {
        location.reload()
    })
}

const originalFetch = window.fetch;
const originalEventSource = window.EventSource

window.fetch = function(url, ...args) {
    if (!url.startsWith('/.assets') && !url.includes('://')) {
        url = `http://localhost:3080${url}`
    }
    console.log('requesting', url)
    return originalFetch(url, ...args);
}
window.EventSource = function(url, ...args) {
    if (!url.startsWith('/.assets') && !url.includes('://')) {
        url = `http://localhost:3080${url}`
    }
    console.log('requesting', url)
    return new originalEventSource(url, ...args)
}
