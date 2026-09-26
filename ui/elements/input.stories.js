export default {
    title: "Elements/Input",
    parameters: {
        docs: {
            description: {
                component: "A disabled input has a dashed border, a subtle background, muted text and the not-allowed cursor.",
            },
        },
    },
    args: { placeholder: "", value: "", disabled: false },
    render: ({ placeholder, value, disabled }) =>
        `<input class="as-input" aria-label="Example" placeholder="${placeholder}" value="${value}"${disabled ? " disabled" : ""}>`,
};

export const Empty = {};

export const Filled = { args: { value: "alice" } };

export const Disabled = { args: { value: "alice", disabled: true } };

export const Textarea = {
    render: () =>
        `<textarea class="as-input leading-relaxed" rows="6" aria-label="Body">A longer body of text.</textarea>`,
};
