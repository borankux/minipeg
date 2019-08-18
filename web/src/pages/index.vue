<template>
    <div class="index">
      <div class="main-box">
        <input type="file" ref="file" style="display: none;" @change="onFileChange">
        <div class="upload-box">
          <div>
            <div class="input-preview" @click="onUploadClicked">
              <img :src="preview" alt="" class="preview-image">
              <div class="upload-mask" v-show="showMask">
                <div class="hint">单击上传</div>
              </div>
            </div>
            <span>大小:{{size_input}}</span>
          </div>
          <div>
            <div class="output-preview">
              <img :src="remote" alt="" class="preview-image">
            </div>
            <span>大小:{{size_output}}</span>
          </div>
        </div>
        <div style="width: 100%;margin-top: 30px; text-align: center;">
          <button class="btn btn-primary" @click="compress">压缩</button>
          <button class="btn" style="margin-left: 10px;" v-show="showDownload" @click="download">下载</button>
        </div>
      </div>
    </div>
</template>

<script>
  import filesize from 'filesize'
  import {uploadFile} from "../apis/comporess";

  export default {
    name: "index",
    data () {
      return {
        preview:"",
        remote:"",
        showMask: true,
        fileName:'',
        size_input:'0',
        size_output:'0',
        showDownload:false,
      }
    },
    methods: {
      onUploadClicked() {
        this.$refs.file.click();
      },
      onFileChange(e)
      {
        if(e.target.files.length >0 ) {
          let file = e.target.files[0];
          this.fileName = file.name;
          this.size_input = filesize(file.size);
          let reader = new FileReader();
          reader.readAsDataURL(file);
          let that = this;
          reader.onloadend = function(e){
            that.preview = e.target.result;
            that.remote = e.target.result;
            that.size_output = '12KB'
            that.showMask = false;
          }
        }
      },
      compress (){
        uploadFile(this.fileName, this.preview).then((res) => {
          console.log(res.data);
        })
      },
      download()
      {

      }
    }
  }
</script>

<style scoped>
  .main-box {
    background-color: #dadada;
    width: 720px;
    box-sizing: border-box;
    padding: 10px;
    margin: 20px auto;
    border: 3px dashed #a8a8a8;
    border-radius: 3px;
  }
  .upload-box {
    display: flex;
    justify-content: center;;
    width: 700px;
    margin: 20px auto;
    text-align: center;
    padding: 10px;
  }

  .input-preview, .output-preview {
    position: relative;
    width: 300px;
    height: 200px;
    background-color: #bfbfbf;
    margin-left: 30px;
    border-radius: 3px;
  }
  .preview-image {
    height: 200px;
    cursor: pointer;
  }


  .upload-mask {
    cursor: pointer;
    position: absolute;
    top: 0;
    left: 0;
    width: 300px;
    height: 200px;
    background-image: url("../assets/upload.png");
    background-repeat: no-repeat;
    background-position: center;
    background-size: 20%;
  }

  .upload-mask .hint {
    width: 100%;
    position: absolute;
    text-align: center;
    font-size: 0.8em;
    bottom: 10px;
    color: #666;
  }
</style>
