<template>
  <div class="container my-5">
    <div class="row justify-content-center">
      <div class="col-8">
        <router-link :to="{ name: 'datatable.index' }" class="btn btn-primary btn-sm rounded shadow mb-3">back</router-link>

        <div class="card rounded shadow">
          <div class="card-header">Tambah Data</div>
          <div class="card-body">
            <form @submit.prevent="store()">
              <div class="mb-3">
                <label for="exampleFormControlInput1" class="form-label">Harga Premium</label>
                <input type="number" class="form-control" id="exampleFormControlInput1" placeholder="Rp." v-model="harga.premium" />
              </div>
              <div class="mb-3">
                <label for="exampleFormControlInput1" class="form-label">Harga Pertalite</label>
                <input type="number" class="form-control" id="exampleFormControlInput1" placeholder="Rp." v-model="harga.pertalite" />
              </div>
              <div class="mb-3">
                <label for="exampleFormControlInput1" class="form-label">Jumlah Liter</label>
                <input type="number" class="form-control" id="exampleFormControlInput1" placeholder="Liter" v-model="harga.liter" />
              </div>
              <button class="btn btn-outline-primary">Submit</button>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import axios from 'axios';

export default {
  setup() {
    // data binding
    const harga = reactive({
      premium: '',
      pertalite: '',
      liter: '',
    });

    const validation = ref([]);

    const router = useRouter();

    // fungsi untuk menyimpan inputan ke database
    function store() {
      axios
        .post('http://localhost:8080/api/bbm', harga)
        .then(() => {
          router.push({
            name: 'datatable.index',
          });
        })
        .catch((err) => {
          validation.value = err.response.data;
        });
    }

    return {
      harga,
      validation,
      router,
      store,
    };
  },
};
</script>
