const colors = [
    "bg",
    "fg",
    "muted",
    "border",
    "subtle",
    "button",
    "button-fg",
    "button-hover",
    "accent",
    "switch",
    "success",
    "info",
    "warning",
    "danger",
];

export default {
    title: "Foundations/Tokens",
};

export const Colors = {
    render: () => `
<ul class="grid grid-cols-2 gap-4 sm:grid-cols-4">
    ${colors
        .map(
            (name) => `
    <li class="flex flex-col gap-2">
        <div class="h-12 rounded-md border border-border" style="background: var(--color-${name})"></div>
        <code class="text-sm text-muted">${name}</code>
    </li>`,
        )
        .join("")}
</ul>`,
};
