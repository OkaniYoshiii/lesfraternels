export const Guards = {
    isHTMLElement,
    isArrayOfHtmlElement,
}

function isHTMLElement(element: unknown): element is HTMLElement {
    return element instanceof HTMLElement
}

function isArrayOfHtmlElement(elements: unknown[]): elements is HTMLElement[] {
    return elements.every((element) => element instanceof HTMLElement)
}