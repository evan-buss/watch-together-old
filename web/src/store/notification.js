import { writable } from "svelte/store";
import { listen } from "svelte/internal";


function createNotifications() {
  const { subscribe, set, update } = writable([]);

  function timedNotification(newNotif, timeout) {
    update(list => {
      // After the timeout, we filter the item out and update the list
      setTimeout(() => {
        update(list => list.filter(item => item.time !== newNotif.time));
      }, timeout);
      return [...list, newNotif];
    });
  }

  return {
    subscribe,
    addPersistant: (newNotif) => {
      update(list => [...list, { ...newNotif, ...{ time: Date.now() } }]);
    },
    addTimed: (newNotif, timeout) => timedNotification({ ...newNotif, ...{ time: Date.now() } }, timeout),
    remove: (notification) => update(list => list.filter(item => item.time !== notification.time)),
    clear: () => set([])
  };
}

export const notifications = createNotifications();