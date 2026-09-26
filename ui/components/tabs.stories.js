export default {
    title: "Components/Tabs",
    args: { active: "Posts" },
    argTypes: {
        active: { control: "select", options: ["Dashboard", "Posts", "Users"] },
    },
    render: ({ active }) => `
<nav aria-label="Administration" class="as-tabs">
    ${["Dashboard", "Posts", "Users"]
        .map((label) => `<a href="#" class="as-tab"${label === active ? ' aria-current="page"' : ""}>${label}</a>`)
        .join("")}
</nav>`,
};

export const Default = {};
