<template>
  <b-form class="mt-3">
    <b-row>
      <b-row>
        <h4 class="text-secondary">Book Details</h4>
      </b-row>
      <b-col cols="6">
        <b-form-group id="title-name" label="Title" label-for="title-name">
          <b-form-input
            id="title-name"
            type="text"
            placeholder="Title"
            v-model="book.Title"
          ></b-form-input>
        </b-form-group>
      </b-col>
      <b-col cols="6">
        <b-form-group id="author-name" label="Author" label-for="author-name">
          <b-form-input
            id="author-name"
            type="text"
            placeholder="Author"
            v-model="book.Author"
          ></b-form-input>
        </b-form-group>
      </b-col>
    </b-row>
    <b-row class="mt-3">
      <b-col cols="6">
        <b-form-group id="call-number" label="CallNumber" label-for="call-number">
          <b-form-input
            id="call-number"
            type="text"
            placeholder="Call Number"
            v-model="book.Call_Number"
          ></b-form-input>
        </b-form-group>
      </b-col>
    </b-row>
    
    <b-row class="mt-4">
      <b-col cols="3">
        <b-button variant="primary" class="px-5" @click="addNewBook"
          >Add Book</b-button
        >
      </b-col>
      <b-col>
        <b-button variant="warning" @click="triggerClose">Close</b-button>
      </b-col>
    </b-row>
  </b-form>
</template>

<script>
import axios from "axios";

export default {
  name: "CreateBookModal",
  data() {
    return {
      book: {},
    };
  },
  methods: {
    triggerClose() {
      this.$emit("closeCreateModal");
    },
    addNewBook() {
      console.log(this.book);
      axios
        .post("http://localhost:3000/book/", this.book)
        .then((response) => {
          console.log(response.data);
          this.$emit("closeCreateModal");
          this.$emit("reloadDataTable");
          this.$emit("showSuccessAlert");
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
};
</script>