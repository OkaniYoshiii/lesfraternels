import { Slider } from './components/slider.js'

const slider = document.querySelector('[data-component="slider"]')

if(slider instanceof HTMLElement) {
    const error = Slider.handle(slider)

    if(error !== null) {
        console.error(error)
    }
}

