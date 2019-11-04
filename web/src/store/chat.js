import { writable } from "svelte/store";

export const messages = writable([]);

// function createMessageStore() {
//   const { subscribe, update, set } = writable([]);

//   return {
//     subscribe,
//     sendMessage: () => {  },
//   };
// }

// export const messages = createMessageStore();