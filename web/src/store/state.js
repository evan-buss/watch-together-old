import { writable } from "svelte/store";


export const sidebarVisible = writable({});

// User contains information about the current user
function createUser() {
  const { subscribe, set, update } = writable({});

  // Load from localstorage
  if (localStorage.getItem('user') !== null) {
    set(JSON.parse(localStorage.getItem('user')));
  }

  return {
    subscribe,
    login: (user) => {
      set(user);
      window.localStorage.setItem('user', JSON.stringify(user))
    },
    logout: () => {
      set({});
      window.localStorage.removeItem('user');
    }
  };
}

export const user = createUser();