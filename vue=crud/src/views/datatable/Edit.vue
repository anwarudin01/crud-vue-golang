<template>
  <div class="container my-5">
    <div class="row justify-content-center">
      <div class="col-8">
        <router-link :to="{ name: 'datatable.index' }" class="btn btn-primary btn-sm rounded shadow mb-3">back</router-link>

        <div class="card rounded shadow">
          <div class="card-header">Edit Data</div>
          <div class="card-body">
            <form @submit.prevent="update()">
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
import { reactive, ref, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import axios from 'axios';

export default {
  setup() {
    // data binding
    let harga = reactive({
      premium: '',
      pertalite: '',
      liter: '',
    });

    const validation = ref([]);

    const router = useRouter();

    // dimanfaatkan untuk mendapatkan data yang ada di parameter url/uri
    const route = useRoute();

    // pada saat componen sudah di tampilkan
    onMounted(() => {
      axios
        .get(`http://localhost:8080/api/bbm/${route.params.id}`)
        .then((result) => {
          harga.premium = result.data.premium;
          harga.pertalite = result.data.pertalite;
          harga.liter = result.data.liter;
        })
        .catch((err) => {
          console.log(err.response.data);
        });
    });

    // fungsi untuk menyimpan inputan ke database
    function update() {
      axios
        .put(`http://localhost:8080/api/bbm/${route.params.id}`, harga)
        .then(() => {
          router.push({
            name: 'datatable.index',
          });
        })
        .catch((err) => {
          console.log(err.response);
        });
    }

    return {
      harga,
      validation,
      router,
      update,
    };
  },
};
</script>
