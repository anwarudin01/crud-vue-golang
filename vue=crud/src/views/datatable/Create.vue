<template>
  <div class="container my-5">
    <div class="row justify-content-center">
      <div class="col-8">
        <router-link :to="{ name: 'datatable.index' }" class="btn btn-primary btn-sm rounded shadow mb-3">back</router-link>

        <div class="card rounded shadow">
          <div class="card-header">Tambah Data</div>
          <div class="card-body">
            <div class="mb-3">
              <label for="exampleFormControlInput1" class="form-label">Harga Premium</label>
              <input type="email" class="form-control" id="exampleFormControlInput1" placeholder="Rp." />
            </div>
            <div class="mb-3">
              <label for="exampleFormControlInput1" class="form-label">Harga Pertalite</label>
              <input type="email" class="form-control" id="exampleFormControlInput1" placeholder="Rp." />
            </div>
            <button class="btn btn-outline-primary">Submit</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {};
</script>
