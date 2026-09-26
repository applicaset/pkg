import arrowLeft from "../../../web/templates/icons/mdi-arrow-left.svg?raw";

export default {
    title: "Components/Back link",
    args: { label: "Back to the blog" },
    render: ({ label }) => `<p><a href="#" class="as-back-link">${arrowLeft}${label}</a></p>`,
};

export const Default = {};
