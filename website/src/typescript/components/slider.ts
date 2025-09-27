import { Guards } from '../lib/guards.js'

export const Slider = {
    handle,
}

function handle(slider: HTMLElement): Error | null {
    const slidesContainer = slider.querySelector('[data-slider-element="slides-container"]')
    const controlsContainer = slider.querySelector('[data-slider-element="controls-container"]')
    const slides = Array.from(slider.querySelectorAll('[data-slider-element="slide"]'))
    const controls = Array.from(slider.querySelectorAll('[data-slider-element="control"]'))

    if(!Guards.isHTMLElement(slidesContainer)) {
        return new Error('Slider: slidesContainer is expected to be an instance of HTMLElement but is something else.')
    }

    if(!Guards.isHTMLElement(controlsContainer)) {
        return new Error('Slider: controls are expected to be an array of HTMLElement but is something else.')
    }

    if(!Guards.isArrayOfHtmlElement(slides)) {
        return new Error('Slider: slides are expected to be an array of HTMLElement but is something else.')
    }

    if(!Guards.isArrayOfHtmlElement(controls)) {
        return new Error('Slider: controls are expected to be an array of HTMLElement but is something else.')
    }

    const control = getActiveControl(controls, slides, slidesContainer)

    console.log(control)
    if(!Guards.isHTMLElement(control)) {
        return new Error('Slider: cannot determine the initial active control.')
    }

    control.style.backgroundColor = 'red'

    slidesContainer.style.transition = 'translate ease-in-out 500ms'

    controlsContainer.addEventListener('click', function(ev) {
        const target = ev.target
        const closestControl = (Guards.isHTMLElement(target)) ? target.closest('[data-slider-element="control"]') : null
        const index = controls.findIndex((control) => control === closestControl)
        const translateInPercent = 100 * index * - 1

        if(Guards.isHTMLElement(closestControl) && index !== -1) {
            controls.forEach((control) => {
                control.style.backgroundColor = ""
            })

            closestControl.style.backgroundColor = 'red'
            slidesContainer.style.translate = `${translateInPercent}%`
        }
    })

    return null
}

function getActiveControl(controls: HTMLElement[], slides: HTMLElement[], slidesContainer: HTMLElement): HTMLElement|undefined {
    const index = getActiveSlideIndex(slides, slidesContainer)
    console.log(slidesContainer.style)

    return index !== -1 ? controls[index] : undefined
}

function getActiveSlideIndex(slides: HTMLElement[], slidesContainer: HTMLElement): number {
    return slides.findIndex((_, index) => slidesContainer.style.translate === index * 100 + "%")
}