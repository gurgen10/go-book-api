<template>
  <div>
    <b-row>
      <b-alert v-model="showSuccessAlert" variant="success" dismissible>
        {{ alertMessage }}
      </b-alert>
    </b-row>
    <b-row>
      <book-overview
        :totalBooks="numberOfBooks"
        @totalBooksIsActive="setFilterTotalIsActive"
      ></book-overview>
    </b-row>
    <b-row class="mt-3">
      <b-card>
        <b-row align-h="between">
          <b-col cols="6">
            <h3>{{ tableHeader }}</h3>
          </b-col>
          <b-col cols="2">
            <b-row>
              <b-col>
                <b-button
                  variant="primary"
                  id="show-btn"
                  @click="showCreateModal"
                >
                  <b-icon-plus class="text-white"></b-icon-plus>
                  <span class="h6 text-white">New Book</span>
                </b-button>
              </b-col>
            </b-row>
          </b-col>
        </b-row>
        <b-row class="mt-3">
          <b-table
            striped
            hover
            :items="items"
            :fields="fields"
            class="text-center"
          >
            <template #cell(title)="data">
              {{
                `${data.item.Title}`
              }}
            </template>
             <template #cell(author)="data">
              {{
                `${data.item.Author}`
              }}
            </template>
            <template #cell(call_number)="data">
              {{
                `${data.item.CallNumber}`
              }}
            </template>
            <template #cell(actions)="data">
              <b-row>
                <b-col cols="7">
                  <b-icon-pencil-square
                    class="action-item"
                    variant="primary"
                    @click="getRowData(data.item.ID)"
                  ></b-icon-pencil-square>
                </b-col>
                <b-col cols="1">
                  <b-icon-trash-fill
                    class="action-item"
                    variant="danger"
                    @click="showDeleteModal(data.item.ID)"
                  ></b-icon-trash-fill>
                </b-col>
              </b-row>
            </template>
          </b-table>
        </b-row>
      </b-card>
    </b-row>

    <!-- Modal for adding new Books -->
    <b-modal
      ref="create-book-modal"
      size="xl"
      hide-footer
      title="New Book"
    >
      <create-book-form
        @closeCreateModal="closeCreateModal"
        @reloadDataTable="getBookData"
        @showSuccessAlert="showAlertCreate"
      ></create-book-form>
    </b-modal>

    <!-- Modal for updating Books -->
    <b-modal
      ref="edit-book-modal"
      size="xl"
      hide-footer
      title="Edit Book"
    >
      <edit-book-form
        @closeEditModal="closeEditModal"
        @reloadDataTable="getBookData"
        @showSuccessAlert="showAlertUpdate"
        :bookId="bookId"
      ></edit-book-form>
    </b-modal>

    <!-- Delete Book Modal -->
    <b-modal
      ref="delete-book-modal"
      size="md"
      hide-footer
      title="Confirm Deletion"
    >
      <delete-book-modal
        @closeDeleteModal="closeDeleteModal"
        @reloadDataTable="getBookData"
        @showDeleteAlert="showDeleteSuccessModal"
        :bookId="bookId"
      ></delete-book-modal>
    </b-modal>
  </div>
</template>

<script>
import axios from "axios";
import BookOverview from "@/components/BookOverview.vue";
import CreateBookForm from "@/components/CreateBookForm.vue";
import EditBookForm from "@/components/EditBookForm.vue";
import DeleteBookModal from "@/components/DeleteBookModal.vue";

export default {
  components: {
    BookOverview,
    CreateBookForm,
    EditBookForm,
    DeleteBookModal,
  },
  data() {
    return {
      // Note 'isActive' is left out and will not appear in the rendered table

      fields: [
        {
          key: "title",
          label: "Title",
          sortable: false,
        },
        {
          key: "author",
          label: "Author",
          sortable: false,
        },
        {
          key: "call_number",
          label: "Call Number",
          sortable: false,
        },
        "actions",
      ],
      items: [],
      numberOfBooks: 0,
      bookId: 0,
      companySearchTerm: "",
      tableHeader: "",
      showSuccessAlert: false,
      alertMessage: "",
    };
  },
  mounted() {
    this.getBookData();
  },
  methods: {
    showCreateModal() {
      this.$refs["create-book-modal"].show();
    },
    closeCreateModal() {
      this.$refs["create-book-modal"].hide();
    },
    getBookData() {
      const url = "http://localhost:3000/books"
      axios
        .get(url,{
            headers: {
              "Access-Control-Allow-Origin": "*",
              "Access-Control-Allow-Credentials": "true",
              "Access-Control-Max-Age": "1800",
              "Access-Control-Allow-Headers": "content-type",
              "Access-Control-Allow-Methods": "PUT, POST, GET, DELETE, PATCH, OPTIONS"
            }
          })
        .then((response) => {
          this.tableHeader = "Total Book";
          this.items = response.data.data;
          console.log(this.items, "getBookData");
          this.numberOfBooks = response.data.data.length;
        })
        .catch((error) => {
          console.log(error);
        });
    },
    getRowData(id) {
      this.$refs["edit-book-modal"].show();
      this.bookId = id;
    },
    closeEditModal() {
      this.$refs["edit-book-modal"].hide();
    },
    setFilterTotalIsActive() {
      this.tableHeader = "Total Books";
      this.getBookData();
    },
    showAlertCreate() {
      this.showSuccessAlert = true;
      this.alertMessage = "Book was created successfully!";
    },
    showAlertUpdate() {
      this.showSuccessAlert = true;
      this.alertMessage = "Book was updated successfully";
    },
    showDeleteModal(id) {
      this.$refs["delete-book-modal"].show();
      this.bookId = id;
    },
    closeDeleteModal() {
      this.$refs["delete-book-modal"].hide();
    },
    showDeleteSuccessModal() {
      this.showSuccessAlert = true;
      this.alertMessage = "Book was deleted successfully!";
    },
  },
};
</script>

<style>
.action-item:hover {
  cursor: pointer;
}
</style>