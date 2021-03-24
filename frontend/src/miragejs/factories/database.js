import { Factory } from "miragejs";
import faker from "faker";

export default {
  database: Factory.extend({
    name(i) {
      // i + 1 to accout for the "*" creation in the instance
      return "db" + faker.fake("{{lorem.word}}") + (i + 1);
    },
    createdTs(i) {
      return Date.now() - (i + 1) * 1800 * 1000;
    },
    lastUpdatedTs(i) {
      return Date.now() - i * 3600 * 1000;
    },
    ownerId() {
      return "100";
    },
    syncStatus(i) {
      if (i % 3 == 0) {
        return "OK";
      }
      if (i % 3 == 1) {
        return "MISMATCH";
      }
      if (i % 3 == 2) {
        return "NOT_FOUND";
      }
    },
    lastSuccessfulSyncTs(i) {
      return Date.now() - i * 3600 * 1000;
    },
    fingerprint(i) {
      return faker.fake("{{random.alpha}}");
    },
  }),
};
