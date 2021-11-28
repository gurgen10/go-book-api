<template>
  <b-row class="mt-4">
    <b-col cols="6">
      <b-button variant="primary" class="px-5" @click="deleteBook"
        >Delete Book</b-button
      >
    </b-col >
    <b-col cols="3">
      <b-button variant="warning" @click="triggerClose">Close</b-button>
    </b-col>
  </b-row>
</template>

<script>
import axios from "axios";

export default {
  name: "DeletBookModal",
  props:{
    bookId: Number 
  },
  methods: {
    triggerClose() {
      this.$emit("closeDeleteModal");
    },
    deleteBook() {
      axios
        .delete(`http://localhost:3000/book/${this.bookId}`)
        .then((response) => {
          console.log(response.data);
          this.$emit("closeDeleteModal");
          this.$emit("reloadDataTable");
          this.$emit("showDeleteAlert");
        })
        .catch((error) => {
          console.log(error);
        });

    }
  }
}
</script>