export default {
    title: "Elements/Check",
    args: { label: "Editor", checked: false },
    render: ({ label, checked }) =>
        `<label class="as-check"><input type="checkbox"${checked ? " checked" : ""}> ${label}</label>`,
};

export const Unchecked = {};

export const Checked = { args: { checked: true } };
