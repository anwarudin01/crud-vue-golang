<template>
  <div class="container my-5">
    <div class="row justify-content-center">
      <div class="col-8">
        <router-link :to="{ name: 'datatable.create' }" class="btn btn-primary btn-sm rounded shadow mb-3">Add</router-link>

        <div class="card rounded shadow">
          <div class="card-header">List Harga BBM</div>
          <table class="table">
            <thead>
              <tr>
                <th scope="col">Jumlah Liter</th>
                <th scope="col">Premium</th>
                <th scope="col">Pertalite</th>
                <th scope="col">Action</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(data, index) in datatable" :key="index">
                <th scope="row">{{ index + 1 }}</th>
                <td>{{ data.premium }}</td>
                <td>{{ data.pertalite }}</td>
                <td>
                  <div class="btn-group">
                    <router-link :to="{ name: 'datatable.edit', params: { id: data.id } }" class="btn btn-sm btn-outline-info me-3">Edit</router-link>
                    <button class="btn btn-sm btn-outline-danger">Delete</button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

// onMounted adalah event yang dilakukan setelah template di eksekusi
// ref untuk menggunakan API compotition
import { onMounted, ref } from 'vue';

export default {
  setup() {
    // definisikan wadah untuk menampung data yang akan di panggil dari API
    // reactive state
    let datatable = ref([]);

    // Apa yang akan dijalankan
    onMounted(() => {
      // get data dari API endpoint
      axios
        .get('http://localhost:8080/api/bbm')
        .then((result) => {
          datatable.value = result.data;
        })
        .catch((err) => {
          console.log(err.response);
        });
    });
    return {
      datatable,
    };
  },
};
</script>
