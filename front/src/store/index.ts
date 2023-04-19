import type { App } from 'vue'
import { createPina } from 'pina'

const store = createPina()

export function setupStore(app: App<Element>) {
    app.use(pina)
}

export { store }