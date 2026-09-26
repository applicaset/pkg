export default {
    title: "Components/Card",
    parameters: {
        docs: {
            description: {
                component: "Only around a figure, never around a whole section.",
            },
        },
    },
};

export const Figures = {
    render: () => `
<dl class="grid grid-cols-2 gap-4 sm:grid-cols-4">
    ${[
        ["Published", 12],
        ["Drafts", 3],
        ["Archived", 1],
        ["Users", 5],
    ]
        .map(
            ([label, count]) => `
    <div class="as-card">
        <dt class="text-sm text-muted">${label}</dt>
        <dd class="mt-1 text-2xl font-semibold">${count}</dd>
    </div>`,
        )
        .join("")}
</dl>`,
};
