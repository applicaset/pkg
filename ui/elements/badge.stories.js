const variants = { neutral: "", success: "variant-success", info: "variant-info", warning: "variant-warning", danger: "variant-danger" };

export default {
    title: "Elements/Badge",
    args: { variant: "neutral", text: "archived" },
    argTypes: {
        variant: { control: "select", options: Object.keys(variants) },
    },
    parameters: {
        docs: {
            description: {
                component: "Post status: published is success, draft is info, archived is neutral.",
            },
        },
    },
    render: ({ variant, text }) => `<span class="as-badge ${variants[variant]}">${text}</span>`,
};

export const Neutral = {};

export const Success = { args: { variant: "success", text: "published" } };

export const Info = { args: { variant: "info", text: "draft" } };

export const Warning = { args: { variant: "warning", text: "warning" } };

export const Danger = { args: { variant: "danger", text: "danger" } };
