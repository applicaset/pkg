export default {
    title: "Foundations/Base",
};

export const Typography = {
    render: () => `
<div class="flex flex-col gap-4">
    <h1>Heading one</h1>
    <h2>Heading two</h2>
    <p>Body text in the system font. <a href="#">A link</a> is underlined with the border color.</p>
    <p class="font-mono text-sm">Monospace text</p>
</div>`,
};
