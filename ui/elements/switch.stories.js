export default {
    title: "Elements/Switch",
    args: { checked: false },
    parameters: {
        docs: {
            description: {
                component:
                    "Only for settings that apply the moment they are flipped. A form with a submit button uses checkboxes.",
            },
        },
    },
    render: ({ checked }) =>
        `<input type="checkbox" role="switch" class="as-switch" aria-label="Open registration"${checked ? " checked" : ""}>`,
};

export const Off = {};

export const On = { args: { checked: true } };
