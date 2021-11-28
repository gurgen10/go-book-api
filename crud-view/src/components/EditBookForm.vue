<template>
  <b-form class="mt-3">
    <b-row>
      <b-row>
        <h4 class="text-secondary">Book Details</h4>
      </b-row>
      <b-col cols="6">
        <b-form-group id="title-name" label="Title Name" label-for="title-name">
          <b-form-input
            id="title-name"
            type="text"
            placeholder="Title Name"
            v-model="book.Title"
          ></b-form-input>
        </b-form-group>
      </b-col>
      <b-col cols="6">
        <b-form-group id="author-name" label="Author Name" label-for="author-name">
          <b-form-input
            id="author-name"
            type="text"
            placeholder="Author Name"
            v-model="book.Author"
          ></b-form-input>
        </b-form-group>
      </b-col>
    </b-row>
    <b-row class="mt-3">
      <b-col cols="6">
         <b-form-group id="call-number" label="Call Number" label-for="call_number">
          <b-form-input
            id="call-number"
            type="text"
            placeholder="Call Number"
            v-model="book.CallNumber"
          ></b-form-input>
        </b-form-group>
      </b-col>
    </b-row>
    
    <b-row class="mt-4">
      <b-col cols="3">
        <b-button variant="primary" class="px-5" @click="updateBook"
          >Update Book</b-button
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
  props: {
    bookId: Number,
  },
  data() {
    return {
      book: {},
    };
  },
  mounted() {
    this.getBookByID();
  },
  methods: {
    triggerClose() {
      this.$emit("closeEditModal");
    },
    getBookByID() {
      axios
        .get(`http://localhost:3000/book/${this.bookId}`)
        .then((response) => {
          console.log(response.data);
          this.book = response.data.data;
          console.log(this.book, "this.book");
        })
        .catch((error) => {
          console.log(error);
        });
    },
    updateBook() {
      console.log(this.book, 'start');
      axios
        .patch(
          `http://localhost:3000/book/${this.bookId}`,
          this.book
        )
        .then((response) => {
          console.log(response.data, "response.data");
          this.$emit("closeEditModal");
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
