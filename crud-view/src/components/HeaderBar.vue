<template>
  <b-navbar toggleable="lg" type="dark" variant="primary">
    <b-container>
      <b-row class="text-white">
        <b-col>
          <b-iconstack font-scale="5" rotate="90">
            <b-icon stacked icon="chevron-right" shift-h="-4" variant="danger"></b-icon>
            <b-icon stacked icon="chevron-right" shift-h="0" variant="success"></b-icon>
            <b-icon stacked icon="chevron-right" shift-h="4" variant="info"></b-icon>
            </b-iconstack>
        </b-col>
        <b-col>
          <span class="h3">Retail Merchant Group</span>
        </b-col>
      </b-row>
    </b-container>
  </b-navbar>
</template>

<script>
export default {
  name: "HeaderBar",
};
</script>