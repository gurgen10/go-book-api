<template>
  <div>
    <b-row class="text-center">
      <b-col cols="3 m-auto">
        <div
          class="filter-card p-3"
        >
          <h4>Total Books: <strong  class="text-success"> {{ totalBooks }}</strong></h4>
        </div>
      </b-col>
    </b-row>
  </div>
</template>

<script>
export default {
  name: "BookOverview",
  props: {
    totalBooks: Number,
  },
};
</script>
