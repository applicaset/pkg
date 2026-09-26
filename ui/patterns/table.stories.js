const posts = [
    ["On Computing", "published", "2026-09-20"],
    ["Notes on Go", "draft", "2026-09-18"],
    ["Old news", "archived", "2025-01-02"],
];

const badges = { published: "variant-success", draft: "variant-info", archived: "" };

export default {
    title: "Patterns/Table",
    parameters: {
        docs: {
            description: {
                component: "Wrap it in `overflow-x-auto`. Row actions are `.as-button.variant-outlined.size-sm`.",
            },
        },
    },
};

export const Posts = {
    render: () => `
<div class="overflow-x-auto">
    <table class="as-table">
        <thead>
            <tr><th>Title</th><th>Status</th><th>Updated</th><th><span class="sr-only">Actions</span></th></tr>
        </thead>
        <tbody>
            ${posts
                .map(
                    ([title, status, date]) => `
            <tr>
                <td>${title}</td>
                <td><span class="as-badge ${badges[status]}">${status}</span></td>
                <td class="whitespace-nowrap text-muted">${date}</td>
                <td class="text-right"><a href="#" class="as-button variant-outlined size-sm">Edit</a></td>
            </tr>`,
                )
                .join("")}
        </tbody>
    </table>
</div>`,
};
