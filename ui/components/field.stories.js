export default {
    title: "Components/Field",
    args: { label: "Username", type: "text", hint: "" },
    argTypes: {
        type: { control: "select", options: ["text", "password", "email"] },
    },
    render: ({ label, type, hint }) => `
<div class="flex max-w-sm flex-col gap-5">
    <label class="as-field">${label} <input class="as-input" type="${type}"></label>
    ${hint ? `<p class="as-hint -mt-3">${hint}</p>` : ""}
</div>`,
};

export const Default = {};

export const WithHint = { args: { label: "Password", type: "password", hint: "At least 8 characters." } };

export const Form = {
    render: () => `
<form class="flex max-w-sm flex-col gap-5">
    <label class="as-field">Username <input class="as-input" name="username" autocomplete="username"></label>
    <label class="as-field">Display name <input class="as-input" name="name" autocomplete="name"></label>
    <label class="as-field">Password <input class="as-input" name="password" type="password" autocomplete="new-password"></label>
    <button type="button" class="as-button w-full">Create account</button>
</form>`,
};
