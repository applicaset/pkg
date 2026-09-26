import plus from "../../../web/templates/icons/mdi-plus.svg?raw";

const variants = { default: "", outlined: "variant-outlined" };
const sizes = { sm: "size-sm", md: "size-md" };

export default {
    title: "Elements/Button",
    args: { label: "Save", variant: "default", size: "md", danger: false, disabled: false, icon: false },
    argTypes: {
        variant: { control: "select", options: Object.keys(variants) },
        size: { control: "select", options: Object.keys(sizes) },
    },
    parameters: {
        docs: {
            description: {
                component:
                    "`.as-button` alone is the default variant at size md. Table row actions are `.as-button.variant-outlined.size-sm`. A create button starts with the plus icon.",
            },
        },
    },
    render: ({ label, variant, size, danger, disabled, icon }) => {
        const classes = ["as-button", variants[variant], sizes[size], danger ? "is-danger" : ""]
            .filter(Boolean)
            .join(" ");

        return `<button type="button" class="${classes}"${disabled ? " disabled" : ""}>${icon ? plus : ""}${label}</button>`;
    },
};

export const Default = {};

export const Outlined = { args: { variant: "outlined", label: "Cancel" } };

export const Danger = { args: { danger: true, label: "Delete" } };

export const OutlinedDanger = { args: { variant: "outlined", danger: true, label: "Delete" } };

export const Small = { args: { variant: "outlined", size: "sm", label: "Edit" } };

export const WithIcon = { args: { icon: true, label: "New post" } };

export const Disabled = { args: { disabled: true } };
