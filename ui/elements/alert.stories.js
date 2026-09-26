const roles = { success: "status", info: "status", warning: "status", danger: "alert" };

export default {
    title: "Elements/Alert",
    args: { variant: "success", text: "Saved." },
    argTypes: {
        variant: { control: "select", options: Object.keys(roles) },
    },
    render: ({ variant, text }) =>
        `<p role="${roles[variant]}" class="as-alert variant-${variant}">${text}</p>`,
};

export const Success = {};

export const Info = { args: { variant: "info", text: "This is your own account." } };

export const Warning = { args: { variant: "warning", text: "This post is not published." } };

export const Danger = { args: { variant: "danger", text: "Wrong username or password." } };
