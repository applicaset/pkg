import arrowLeft from "../../../web/templates/icons/mdi-arrow-left.svg?raw";
import plus from "../../../web/templates/icons/mdi-plus.svg?raw";

export default {
    title: "Elements/Icon",
    parameters: {
        docs: {
            description: {
                component:
                    "MDI SVGs from each service's `templates/icons/`, sized by `.as-icon` and colored by `currentColor`.",
            },
        },
    },
};

export const All = {
    render: () => `<div class="flex gap-4">${arrowLeft}${plus}</div>`,
};
