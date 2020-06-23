<template>
  <div class="container">
    <el-container class="elcontainer">
      <el-header height="80px" :style="{'backgroundColor':this.navColor}">
        <div class="logo">
          <img v-if="logoUrlLogin" :src="logoUrlLogin" alt />
        </div>
        <div class="title">{{webTitle}}管理系统</div>
        <!-- <div class="rightName">{{userInfo.user_name}}</div> -->
        <el-dropdown trigger="click">
          <span class="el-dropdown-link">
            {{userInfo.user_name}}
            <i class="el-icon-arrow-down el-icon--right"></i>
            <a class="userHead" v-if="!isbeijing2022">
              <img :src="userInfo.img_path" alt />
            </a>
          </span>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item>
              <span @click="signOut()">退出登陆</span>
            </el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </el-header>
      <el-container>
        <el-aside width="130px">
          <p class="aside">| | |</p>
          <el-menu
            default-active="1-1"
            class="el-menu-vertical-demo"
            background-color="#1d2b36"
            text-color="#ffffff"
            @select="handelSelect"
          >
            <el-submenu index="1">
              <template slot="title">
                <i class="el-icon-location"></i>
                <span class="nav_span">页面管理</span>
              </template>
              <el-menu-item-group>
                <el-menu-item
                  v-for="(item,index) in ChannelListInfo"
                  :index="'1-'+item.id"
                  :key="index"
                  style="paddingLeft:55px"
                  @click="setCurPage(item, index)"
                >
                  <span>
                    <i
                      :style="{visibility: index>1?'visible':'hidden'}"
                      class="el-icon-sort-up"
                      @click.stop="sortChannel('up', index)"
                    ></i>
                    <i
                      :style="{visibility: index>0&&index< ChannelListInfo.length - 1?'visible':'hidden'}"
                      class="el-icon-sort-down"
                      @click.stop="sortChannel('down', index)"
                    ></i>
                  </span>
                  <span>{{item.page_name}}</span>
                  <span>
                    <i
                      v-if="index>0"
                      class="el-icon-delete"
                      @click.stop="deleteChannel(item, index)"
                    ></i>
                  </span>
                </el-menu-item>
              </el-menu-item-group>
            </el-submenu>
            <el-menu-item index="2" style="paddingLeft:0px" @click="addChannel">
              <i class="el-icon-menu"></i>
              <span slot="title" class="nav_span">新增频道</span>
            </el-menu-item>
            <el-submenu index="3">
              <template slot="title">
                <i class="el-icon-location"></i>
                <span class="nav_span">基本设置</span>
              </template>
              <el-menu-item-group>
                <el-menu-item index="3" @click="set111">
                  <span
                    slot="title"
                    class="nav_span"
                    style="width:120px;vertical-align: middle;line-height: 16px;"
                  >
                    <i class="el-icon-document"></i>logo&设置
                  </span>
                </el-menu-item>
                <el-menu-item index="5" @click="set111" class="basical">
                  <!-- <i class="el-icon-sort-down"></i> -->
                  <span
                    slot="title"
                    class="nav_span"
                    style="width:120px;vertical-align: middle;line-height: 16px;"
                  >
                    <i class="el-icon-document"></i>积分配制
                  </span>
                </el-menu-item>
              </el-menu-item-group>
            </el-submenu>
            <el-menu-item index="4" style="paddingLeft:0px">
              <i class="el-icon-document"></i>
              <span slot="title" class="nav_span">底栏设置</span>
            </el-menu-item>
            <el-menu-item index="7" style="paddingLeft:0px">
              <i class="el-icon-document"></i>
              <span slot="title" class="nav_span">反馈列表</span>
            </el-menu-item>
            <!-- <el-menu-item
              v-if="webid==111111"
              index="6"
              style="paddingLeft:0px"
            >
              <i class="el-icon-document"></i>
              <span
                slot="title"
                class="nav_span"
              >开发者加模板</span>
            </el-menu-item>-->
          </el-menu>
        </el-aside>
        <el-container>
          <el-main v-if="!basicSetIsShow&&activeMenuItem!=4&&activeMenuItem!=6&&activeMenuItem!=7">
            <div class="channerName">
              <span class="title_span">频道名称</span>
              <el-input v-model="nowChannelTitle" @change="inChange()"></el-input>
              <el-button @click="setChannerName(0)">确定</el-button>
            </div>
            <div class="modelChoose" v-if="!isAddChannel">
              <span class="title_span">模板选择</span>
              <ul>
                <li
                  v-for="(item,index) in WebListInfo"
                  @click="modelChoose(item.id)"
                  v-show="moduleShow(index,item)"
                  :key="index"
                >
                  <img class="webImg" :src="item.show_tpl_pic" alt />
                  <i>
                    {{item.tpl_name}}
                    <!-- <el-button type="primary" @click.stop="">修改缩略图</el-button> -->
                    <input
                      type="file"
                      accept="image/*"
                      @click.stop
                      @change="setModulImg(item.id, index, $event)"
                    />
                  </i>
                  <img
                    class="chooseIcon"
                    v-if="nowModel == item.id"
                    src="../assets/img/chooseModel.png"
                    alt
                  />
                </li>
              </ul>
            </div>
            <div
              class="setSlideshow"
              v-if="!isAddChannel && nowModel !== 4&&nowModel !== 34&&nowModel!==7&&nowModel!==9&& nowModel !== 8&& nowModel !== 30&& nowModel !== 31&& nowModel !== 32&& nowModel !== 33 && nowModel !== 38 && nowModel !== 39 && nowModel !== 40"
            >
              <h3 class="title_h3">轮播图设置</h3>
              <div class="setNumber">
                <div class="setNumberCon">
                  <div class="number">
                    <span class="title_span">数量</span>
                    <el-input-number
                      controls-position="right"
                      :min="0"
                      :max="10"
                      v-model="slideshowNumber"
                    ></el-input-number>
                  </div>
                  <el-button @click="setSlideshowNumber">确定</el-button>
                </div>
              </div>
              <ul>
                <li v-for="(item,index) in bannerList " :key="index">
                  <img :src="item.photo1" alt />
                  <div class="imgId" style="display:flex;display:-webkit-flex;">
                    <!-- <span>组ID</span> -->
                    <i-d-type v-model="item.body10"></i-d-type>
                    <div v-if="item.body10==3" style="position:relative">
                      <el-button type="primary" size="small">上传</el-button>
                      <input
                        style="position: absolute;bottom: 0;z-index: 99999;opacity: 0;width:100%"
                        type="file"
                        accept="image/*"
                        @change="uploadSpread($event,item,16)"
                      />
                    </div>
                    <el-input id="el-input" v-else v-model="item.body9"></el-input>
                  </div>
                  <el-upload
                    :show-file-list="false"
                    style="margin-top:20px;"
                    class="upload-demo"
                    action
                    :auto-upload="false"
                    :on-change="UploadfileElcarousel"
                  >
                    <el-button size="small" type="primary" @click="uploadCarousel(item)">点击上传</el-button>
                    <div
                      slot="tip"
                      class="el-upload__tip"
                    >请上传 {{nowModel==1 ? '1200x950' : (nowModel == 36 ? '897x504' : (nowModel == 37 ? '1920x480' : (nowModel==2 ? '1200x950' : (nowModel==3 ? '1600×562' :(nowModel==10? '1920x800' : ( nowModel==13 ? '1200x950' : (nowModel==29 ? '1440x450': (nowModel == 44 ? '1920x480' : '1200x1060'))))))))}}或同等比列的图</div>
                  </el-upload>
                </li>
              </ul>
            </div>

            <div class="setSlideshow" v-if="!isAddChannel && nowModel == 44">
              <h3 class="title_h3">编辑推荐位</h3>
              <div class="setNumber">
                <div class="setNumberCon">
                  <div class="number">
                    <span class="title_span">数量</span>
                    <el-input-number
                      controls-position="right"
                      :min="0"
                      :max="10"
                      v-model="nationalNumber"
                    ></el-input-number>
                  </div>
                  <el-button @click="setNationalNumber">确定</el-button>
                </div>
              </div>
              <ul>
                <li v-for="(item,index) in nationalList " :key="index">
                  <img :src="item.photo1" alt />
                  <div class="imgId" style="display:flex;display:-webkit-flex;">
                    <!-- <span>组ID</span> -->
                    <i-d-type v-model="item.body10"></i-d-type>
                    <div v-if="item.body10==3" style="position:relative">
                      <el-button type="primary" size="small">上传</el-button>
                      <input
                        style="position: absolute;bottom: 0;z-index: 99999;opacity: 0;width:100%"
                        type="file"
                        accept="image/*"
                        @change="uploadSpread($event,item,17)"
                      />
                    </div>
                    <el-input id="el-input" v-else v-model="item.body9"></el-input>
                  </div>
                  <el-upload
                    :show-file-list="false"
                    style="margin-top:20px;"
                    class="upload-demo"
                    action
                    :auto-upload="false"
                    :on-change="UploadfileElcarousel"
                  >
                    <el-button size="small" type="primary" @click="uploadCarousel(item)">点击上传</el-button>
                    <div slot="tip" class="el-upload__tip">请上传632x430（第一张）308x173（后四张）同比例的图片</div>
                  </el-upload>
                </li>
              </ul>
            </div>

            <!-- <div class="setSlideshow" v-if="!isAddChannel && nowModel === 6">
              <h3 class="title_h3">未登录轮播图设置</h3>
              <div class="setNumber">
                <span class="title_span">数量</span>
                <el-input-number controls-position="right" :min="1" :max="10" v-model="slideNumber"></el-input-number>
                <el-button @click="setSlideNumber">确定</el-button>
              </div>
              <ul>
                <li v-for="(item,index) in slideList " :key="index">
                  <img :src="item.photo1" alt="">
                  <p class="imgId">
                    <span>组ID</span>
                    <el-input id="el-input" v-model="item.group_id"></el-input>
                  </p>
                  <p class="imgUrl">
                    <span>图片链接</span>
                    <el-input v-model="item.link_url" :disabled="true"></el-input>
                  </p>
                  <p class="warn">
                    <span class="text">请上传 {{nowModel==1 ? '1120×884' : (nowModel==2 ? '1120×884' : (nowModel==3 ? '1600×562' :'1120×884'))}} 或同等比列的图</span>
                    <a :href="item.link_url" target="_blank">
                      <el-button type="primary">预览</el-button>
                    </a>
                  </p>
                </li>
              </ul>
            </div>-->
            <div
              class="setRecommend"
              v-if="!isAddChannel && nowModel !== 4&&nowModel !== 34&& nowModel!==10&&nowModel!==7&&nowModel!==9&& nowModel !== 30&& nowModel !== 31&& nowModel !== 32&&nowModel !== 33 &&nowModel !== 37 && nowModel !== 38 && nowModel !== 39 && nowModel !== 40"
            >
              <h3 class="title_h3">推荐专题设置</h3>
              <div class="setNumber">
                <div class="setNumberCon">
                  <div class="number">
                    <span class="title_span">数量</span>
                    <el-input-number
                      controls-position="right"
                      :min="0"
                      :max="10"
                      v-model="showNumber"
                    ></el-input-number>
                  </div>
                  <!-- <div class="groupid">
                    <span>分类ID</span>
                    <el-input v-model="ztid"></el-input>
                  </div>-->
                  <el-button @click="setShowNumber">确定</el-button>
                </div>
              </div>

              <ul>
                <li v-for="(item,index) in hotSpecialList" :key="index">
                  <img :src="item.photo1" alt />
                  <div class="imgTitle" style="display:flex;display:-webkit-flex">
                    <!-- <span>组ID</span> -->
                    <i-d-type v-model="item.body10"></i-d-type>
                    <div v-if="item.body10==3" style="position:relative">
                      <el-button type="primary" size="small">上传</el-button>
                      <input
                        style="position: absolute;bottom: 0;z-index: 99999;opacity: 0;width:100%"
                        type="file"
                        accept="image/*"
                        @change="uploadSpread($event,item,19)"
                      />
                    </div>
                    <el-input v-else id="el-input" v-model="item.body9"></el-input>
                  </div>

                  <el-upload
                    :show-file-list="false"
                    style="margin-top:20px;"
                    class="upload-demo"
                    action
                    :auto-upload="false"
                    :on-change="UploadfileSpecial"
                  >
                    <el-button size="small" type="primary" @click="uploadSpecial(item,'A')">点击上传</el-button>
                    <div
                      slot="tip"
                      class="el-upload__tip"
                    >请上传 {{nowModel==1 ? '800x500' : (nowModel == 36 ? '251x160' : (nowModel==2 ? '1200x600' : (nowModel==3 ? '800×500' :(nowModel==4 ?'1200×600' :(nowModel==11 ?'1200x1060' :(nowModel==8 ?'800x500' :(nowModel == 44 ? '308x144' : '800x500')))))))}} 或同等比列的图</div>
                  </el-upload>
                  <!-- <p class="warn">请上传 {{nowModel==1 ? '600×480' : (nowModel==2 ? '876×432' : (nowModel==3 ? '800×500' :(nowModel==4 ?'1200×600' :'600×480')))}} 或同等比列的图</p>
                  <el-button type="primary">预览</el-button>-->
                </li>
              </ul>
            </div>
            <div class="setRecommend" v-if="!isAddChannel && nowModel == 8">
              <h3 class="title_h3">专题策划上左大专题设置</h3>
              <div class="setNumber">
                <div class="setNumberCon">
                  <div class="number">
                    <span class="title_span">数量</span>
                    <el-input-number
                      controls-position="right"
                      :min="0"
                      :max="1"
                      v-model="showNumberB"
                    ></el-input-number>
                  </div>
                  <!-- <div class="groupid">
                    <span>分类ID</span>
                    <el-input v-model="ztidB"></el-input>
                  </div>-->
                  <el-button @click="setShowNumberB">确定</el-button>
                </div>
              </div>

              <ul>
                <li v-for="(item,index) in hotSpecialListB" :key="index">
                  <img :src="item.photo1" alt />
                  <div class="imgTitle" style="display:flex;display:-webkit-flex">
                    <!-- <span>组ID</span> -->
                    <i-d-type v-model="item.body10"></i-d-type>
                    <div v-if="item.body10==3" style="position:relative">
                      <el-button type="primary" size="small">上传</el-button>
                      <input
                        style="position: absolute;bottom: 0;z-index: 99999;opacity: 0;width:100%"
                        type="file"
                        accept="image/*"
                        @change="uploadSpread($event,item,20)"
                      />
                    </div>
                    <el-input v-else id="el-input" v-model="item.body9"></el-input>
                  </div>
                  <el-upload
                    :show-file-list="false"
                    style="margin-top:20px;"
                    class="upload-demo"
                    action
                    :auto-upload="false"
                    :on-change="UploadfileSpecial"
                  >
                    <el-button size="small" type="primary" @click="uploadSpecial(item,'B')">点击上传</el-button>
                    <div slot="tip" class="el-upload__tip">请上传900x1000或同等比列的图</div>
                  </el-upload>
                  <!-- <p class="warn">请上传 {{nowModel==1 ? '600×480' : (nowModel==2 ? '876×432' : (nowModel==3 ? '800×500' :(nowModel==4 ?'1200×600' :'600×480')))}} 或同等比列的图</p>
                  <el-button type="primary">预览</el-button>-->
                </li>
              </ul>
            </div>
            <div class="setRecommend" v-if="!isAddChannel && nowModel == 8">
              <h3 class="title_h3">专题策划上右小图专题设置</h3>
              <div class="setNumber">
                <div class="setNumberCon">
                  <div class="number">
                    <span class="title_span">数量</span>
                    <el-input-number
                      controls-position="right"
                      :min="0"
                      :max="4"
                      v-model="showNumberC"
                    ></el-input-number>
                  </div>
                  <!-- <div class="groupid">
                    <span>分类ID</span>
                    <el-input v-model="ztidC"></el-input>
                  </div>-->
                  <el-button @click="setShowNumberC">确定</el-button>
                </div>
              </div>

              <ul>
                <li v-for="(item,index) in hotSpecialListC" :key="index">
                  <img :src="item.photo1" alt />
                  <div class="imgTitle" style="display:flex; display:-webkit-flex">
                    <!-- <span>组ID</span> -->
                    <i-d-type v-model="item.body10"></i-d-type>
                    <div v-if="item.body10==3" style="position:relative">
                      <el-button type="primary" size="small">上传</el-button>
                      <input
                        style="position: absolute;bottom: 0;z-index: 99999;opacity: 0;width:100%"
                        type="file"
                        accept="image/*"
                        @change="uploadSpread($event,item,21)"
                      />
                    </div>
                    <el-input v-else id="el-input" v-model="item.body9"></el-input>
                  </div>
                  <el-upload
                    :show-file-list="false"
                    style="margin-top:20px;"
                    class="upload-demo"
                    action
                    :auto-upload="false"
                    :on-change="UploadfileSpecial"
                  >
                    <el-button size="small" type="primary" @click="uploadSpecial(item,'C')">点击上传</el-button>
                    <div slot="tip" class="el-upload__tip">请上传1200x800或同等比列的图</div>
                  </el-upload>
                  <!-- <p class="warn">请上传 {{nowModel==1 ? '600×480' : (nowModel==2 ? '876×432' : (nowModel==3 ? '800×500' :(nowModel==4 ?'1200×600' :'600×480')))}} 或同等比列的图</p>
                  <el-button type="primary">预览</el-button>-->
                </li>
              </ul>
            </div>

            <div
              class="setSlideshow"
              v-if="!isAddChannel && (nowModel === 34 || nowModel === 38 || nowModel === 39)"
            >
              <h3 class="title_h3">推荐视频设置</h3>
              <div class="setNumber">
                <span class="title_span">数量</span>
                <el-input-number
                  controls-position="right"
                  :min="nowModel === 39 ? 1 : 3"
                  :max="nowModel === 39 ? 1 : 5"
                  v-model="videoNumber"
                ></el-input-number>
                <el-button @click="setVideoNumber">确定</el-button>
              </div>
              <ul>
                <li
                  v-for="(item,index) in videoList "
                  :key="index"
                  style="display:flex;flex-direction:column;justify-content: space-between;"
                >
                  <div style="width: 100%;flex-direction: column;height: 88%;">
                    <video autoplay :src="item.photo1" style="width:100%;height:100%"></video>
                  </div>
                  <div style="display:flex;align-items:center">
                    <span style="margin-right:5px">资源ID:</span>
                    <span style="margin-right:5px">
                      <el-input v-model="item.video_id"></el-input>
                    </span>
                    <el-button type="primary" size="mini" @click="setVideoAssetID(item,index)">预览</el-button>
                  </div>
                  <!-- <img :src="item.photo1" alt=""> -->
                  <!-- <p class="imgId">
                    <span>资源ID</span>
                    <el-input id="el-input" style="width:50%" v-model="item.group_id"></el-input>
                  </p>
                  <p class="warn">
                      <el-button type="primary">预览</el-button>
                  </p>-->
                </li>
              </ul>
            </div>

            <!-- <div class="setRecommend" v-if="!isAddChannel && nowModel === 6">
              <h3 class="title_h3">顶部广告位设置</h3>
              <div class="setNumber">
                <span class="title_span">数量</span>
                <el-input-number controls-position="right" :min="1" :max="1" v-model="topAdvertNumber"></el-input-number>
                <el-button @click="setTopAdvertNumber">确定</el-button>
              </div>
              <ul>
                <li v-for="(item,index) in topAdvertList" :key="index">
                  <img :src="item.link_url" alt="">
                  <p class="warn">请上传最小尺寸为 392/240 的图</p>
                  <el-button type="primary">预览</el-button>
                  <div class="upload">
                    <span class="upload">上传/修改</span>
                    <input type="file" accept="image/*" @change="upTopAdvert($event,item)">
                  </div>
                </li>
              </ul>
            </div>-->
            <div class="setRecommend" v-if="!isAddChannel && nowModel === 44">
              <h3 class="title_h3">广告位设置</h3>
              <div class="setNumber">
                <span class="title_span">数量</span>
                <el-input-number
                  controls-position="right"
                  :min="1"
                  :max="4"
                  v-model="bottomAdvertNumber"
                ></el-input-number>
                <el-button @click="setBottomAdvertNumber">确定</el-button>
              </div>
              <ul>
                <li
                  v-for="(item,index) in bottomAdvertList"
                  :key="index"
                  style="width: 300px;height:300px"
                >
                  <img :src="item.photo1" alt />
                  <p class="warn" style="text-align: center">请上传尺寸为 1280x140 的图</p>
                  <div class="upload">
                    <span class="upload">上传/修改</span>

                    <input type="file" accept="image/*" @change="upBottomAdvert($event,item)" />
                  </div>
                  <div style="clear:both"></div>
                  <div style="margin: 20px 0">
                    <span>当前链接： {{ form.ad_url ? form.ad_url : item.ad_url }}</span>
                    <el-form ref="form" :model="form" label-width="80px">
                      <el-form-item label="修改链接">
                        <el-input
                          v-model="form.ad_url"
                          placeholder="请输入链接"
                          @blur="handleBlurUrl(item)"
                        ></el-input>
                      </el-form-item>
                    </el-form>
                  </div>
                </li>
              </ul>
            </div>
            <div
              class="setTag"
              v-if="!isAddChannel &&nowModel !== 1&&nowModel !== 34&& nowModel !== 4&& nowModel!==10&& nowModel!==8&&nowModel!==7&&nowModel!==9&&nowModel!==29&& nowModel !== 30&& nowModel !== 31&& nowModel !== 32&& nowModel !== 33 && nowModel !== 36 && nowModel !== 38 && nowModel !== 39 && nowModel !== 40 && nowModel !== 44"
            >
              <h3 class="title_h3">标签设置</h3>
              <div class="setNumber">
                <span class="title_span">数量</span>
                <el-input-number controls-position="right" :min="1" :max="16" v-model="tagNumber"></el-input-number>
                <el-button @click="setTagNumber">确定</el-button>
              </div>
              <div class="setTitleOrder">
                <template>
                  <el-table :data="nowLabelInfo" style="width:567px">
                    <el-table-column prop="title" label="标签" width="167">
                      <template slot-scope="scope">
                        <input v-model="scope.row.title" />
                      </template>
                    </el-table-column>
                    <el-table-column prop="category_id" label="分类ID" width="200">
                      <template slot-scope="scope">
                        <input v-model="scope.row.category_id" />
                      </template>
                    </el-table-column>
                    <el-table-column prop="address" label="排序" width="200">
                      <template slot-scope="scope">
                        <span @click="sort_down(scope.$index)">
                          <img src="../assets/img/down.png" alt />
                        </span>
                        <span @click="sort_up(scope.$index)">
                          <img src="../assets/img/up.png" alt />
                        </span>
                        <span></span>
                      </template>
                    </el-table-column>
                  </el-table>
                </template>
              </div>
            </div>

            <div
              class="setShowType"
              v-if="!isAddChannel && (nowModel === 40 || nowModel === 39 ||nowModel === 38 || nowModel === 34||nowModel === 33||nowModel === 4||nowModel == 8||nowModel == 1||nowModel == 3||nowModel == 13||nowModel == 29||nowModel == 32)"
            >
              <h3 class="title_h3">专题</h3>
              <span class="title_span">展示分类编号</span>
              <el-input v-model="specialId"></el-input>
              <el-button @click="setShowType">确定</el-button>
            </div>

            <div class="setKeywords" v-if="!isAddChannel &&nowModel == 44">
              <h3 class="title_h3" style="margin:20px 0px">热搜关键词（Tips: 回车输入关键词）</h3>
              <el-tag
                :key="tag"
                v-for="tag in dynamicTags"
                closable
                :disable-transitions="false"
                @close="handleClose(tag)"
              >{{tag}}</el-tag>
              <el-input
                class="input-new-tag"
                v-if="inputVisible"
                v-model="inputValue"
                ref="saveTagInput"
                size="small"
                @keyup.enter.native="handleInputConfirm"
                @blur="handleInputConfirm"
              ></el-input>
              <el-button v-else class="button-new-tag" size="small" @click="showInput">+ 请输入新的关键词</el-button>
              <el-button class="button-new-tag" type="success" @click="handleSaveKeywords">保存关键词</el-button>

              <el-divider></el-divider>
            </div>

            <div
              class="messageSet"
              v-if="!isAddChannel && (nowModel === 1||nowModel==13||nowModel==44)"
            >
              <h3 class="title_h3">
                通知
                <span class="addMessage" @click="openMe({type: 1})">新建通知</span>
              </h3>
              <ul>
                <li v-for="(item,index) in systemList" v-if="index<4" :key="index">
                  <div class="meLeft" @click="openMe({type: 0,item: item})">
                    <p>{{item.title}}</p>
                    <span>{{item.create_time .slice(0, 10)}}</span>
                  </div>
                  <div class="meRight" @click="deleteMessage(item, index)">删除</div>
                </li>
              </ul>
              <p>
                <span class="moreMessage" @click="moreDialogVisible = true">查看更多</span>
              </p>
              <el-dialog class="setDialog" :visible.sync="moreDialogVisible" width="700px">
                <h3 class="title_h3">
                  通知
                  <span class="addMessage" @click="openMe({type: 1})">新建通知</span>
                </h3>
                <ul>
                  <li v-for="(item,index) in systemList" :key="index">
                    <div class="meLeft" @click="openMe({type: 0,item: item})">
                      <p>{{item.title}}</p>
                      <span>{{item.create_time .slice(0, 10)}}</span>
                    </div>
                    <div class="meRight" @click="deleteMessage(item, index)">删除</div>
                  </li>
                </ul>
              </el-dialog>
              <el-dialog
                class="setDialog"
                :title="setMessage.title"
                :visible.sync="dialogVisible"
                width="600px"
              >
                <p>
                  通知标题：
                  <el-input placeholder="请输入通知标题" v-model="setMessage.item.title"></el-input>
                </p>
                <p>
                  通知内容：
                  <el-input type="textarea" placeholder="请输入通知内容" v-model="setMessage.item.body"></el-input>
                </p>
                <span slot="footer" class="dialog-footer">
                  <el-button @click="dialogVisible = false">取 消</el-button>
                  <el-button type="primary" @click="submeMessage">确 定</el-button>
                </span>
              </el-dialog>
            </div>
            <div class="setBtn" v-if="!isAddChannel && nowModel !== 4">
              <el-button type="danger" @click="removeAll">取消所有设置</el-button>
              <el-button type="primary" @click="setAll">提交所有设置</el-button>
            </div>
          </el-main>
          <el-main v-if="basicSetIsShow&&activeMenuItem==3" class="basicSet">
            <div style="margin-bottom:35px;">
              <div style="margin-bottom:15px;">
                <span style="margin-right:35px">频道导航菜单位置</span>
                <el-button size="mini" @click="menuPosSet(1)">居左</el-button>
                <el-button size="mini" @click="menuPosSet(2)">居中</el-button>
                <el-button size="mini" @click="menuPosSet(3)">居右</el-button>
              </div>
            </div>
            <div class="block">
              <span class="demonstration">导航栏底色</span>
              <el-color-picker v-model="navColor"></el-color-picker>
            </div>
            <div class="block">
              <span class="demonstration">网站底栏底色</span>
              <el-color-picker v-model="navColorBot"></el-color-picker>
            </div>

            <div class="block">
              <span class="demonstration">网站slogen底色</span>
              <el-color-picker v-model="navColorSlogen"></el-color-picker>
            </div>
            <div class="block">
              <span class="demonstration">网站slogen</span>
              <el-input class="slogen" v-model="slogentext"></el-input>
              <el-button type="danger" @click="setslogentext">确定</el-button>
            </div>
            <div class="setLogo">
              <span class="demonstration">头部Logo</span>
              <img class="logoPic" :src="logoUrl" />
              <p class="uploadText">尺寸为120×70</p>
              <div class="upload">
                <span class="upload">上传/修改</span>
                <input type="file" accept="image/*" @change="uploadLogo" />
              </div>
            </div>
            <div class="setLogo">
              <span class="demonstration">登录Logo</span>
              <img class="logoPic" :src="logoUrlLogin" />
              <p class="uploadText">尺寸为120×70</p>
              <div class="upload">
                <span class="upload">上传/修改</span>
                <input type="file" accept="image/*" @change="uploadLogoLogin" />
              </div>
            </div>
            <div class="setLogo">
              <span class="demonstration">底部Logo</span>
              <img class="logoPic" :src="logoUrlBot" />
              <p class="uploadText">尺寸为120×70</p>
              <div class="upload">
                <span class="upload">上传/修改</span>
                <input type="file" accept="image/*" @change="uploadLogoBot" />
              </div>
            </div>

            <!-- <div class="setTpl">
              <span class="demonstration">模板</span>
              <img class="tplPic" :src="tplUrl"></img>
              <el-input v-model="tplId" placeholder="模板号"></el-input>
              <div class="uploadTpl">
                <span class="upload">上传/修改</span>
                <input type="file" accept="image/*" @change="uploadTpl">
              </div>
            </div>-->
            <div class="setBtn">
              <!-- <el-button type="primary" @click="setTpl">提交模板设置</el-button> -->
              <el-button type="primary" @click="setColor">提交颜色设置</el-button>
              <el-button type="primary" @click="setLogo">提交logo设置</el-button>
              <el-button type="primary" @click="removelogo">取消logo设置</el-button>
            </div>
            <h3 class="classificationTitle">登录设置</h3>
            <div class="block">
              <el-checkbox v-model="islogin" @change="beforeLoginClick">是否先登录后浏览</el-checkbox>
            </div>
            <h3 class="classificationTitle">侧边栏</h3>
            <div class="block">
              <el-checkbox v-model="isShow_sidebar" @change="showSidebarClick">是否显示侧边栏</el-checkbox>
            </div>
            <div class="block" v-if="isShow_sidebar">
              <span class="demonstration">客服QQ号</span>
              <el-input class="slogen" v-model="sidebarQQ"></el-input>
            </div>
            <div class="block" v-if="isShow_sidebar">
              <span class="demonstration">客服电话</span>
              <el-input class="slogen" v-model="sidebarTel"></el-input>
            </div>
            <div class="block" v-if="isShow_sidebar">
              <span class="demonstration">工作时间</span>
              <el-input class="slogen" v-model="sidebarTime"></el-input>
              <el-button type="danger" @click="setSidebarData">确定</el-button>
            </div>
            <div class="classification">
              <h3 class="classificationTitle">分类</h3>
              <p class="classificationText">分类显示</p>
              <el-table
                :data="business"
                tooltip-effect="dark"
                :span-method="BusinessSpanMethod"
                border
                stripe
                style="width: 600px"
              >
                <el-table-column prop="title" label="分类类型" align="center">
                  <template slot-scope="scope">
                    <span>业务分类</span>
                  </template>
                </el-table-column>
                <el-table-column prop="cname" label="分类" header-align="center" align="center"></el-table-column>
                <el-table-column prop label="显示设置" align="center" width>
                  <template slot-scope="scope">
                    <el-switch
                      v-model="scope.row.active"
                      active-text="on"
                      inactive-color="#F04134"
                      inactive-text="off"
                      @change="changeSwitch(scope.row)"
                    ></el-switch>
                  </template>
                </el-table-column>
              </el-table>
              <el-table
                :data="base"
                tooltip-effect="dark"
                :span-method="BaseSpanMethod"
                border
                stripe
                style="width: 600px"
              >
                <el-table-column prop="title" label="分类类型" align="center">
                  <template slot-scope="scope">
                    <span>基础分类</span>
                  </template>
                </el-table-column>
                <el-table-column prop="cname" label="分类" header-align="center" align="center"></el-table-column>
                <el-table-column prop label="显示设置" align="center" width>
                  <template slot-scope="scope">
                    <el-switch
                      v-model="scope.row.active"
                      active-text="on"
                      inactive-color="#F04134"
                      inactive-text="off"
                      @change="changeSwitch(scope.row)"
                    ></el-switch>
                  </template>
                </el-table-column>
              </el-table>
            </div>
          </el-main>
          <el-main class="home-footer-setting integralSet" v-if="activeMenuItem==5">
            <div class="block">
              <span class="demonstration">上传资源获得积分:</span>
              <el-input class="slogen" placeholder="请输入内容" v-model="upload_score" type="number"></el-input>
            </div>
            <div class="block">
              <span class="demonstration">购买商业图片获得积分:</span>
              <el-input class="slogen" placeholder="请输入内容" v-model="buy_score" type="number"></el-input>
            </div>
            <div class="block">
              <span class="demonstration">每日登录获得积分:</span>
              <el-input class="slogen" placeholder="请输入内容" v-model="login_score" type="number"></el-input>
            </div>
            <div class="setBtn">
              <el-button type="primary" class="BtnStyle" @click="UniversalUpdata">确认修改</el-button>
            </div>
            <el-table
              :data="AllCategories"
              tooltip-effect="dark"
              :span-method="tableObjSpanMethod"
              border
              stripe
              style="width: 600px"
            >
              <el-table-column prop="cname" label="分类" align="center"></el-table-column>
              <el-table-column prop="score" label="下载资源所需积分" header-align="center" align="center">
                <template slot-scope="scope">
                  <el-input v-model="scope.row.score" placeholder="请输入内容" type="number"></el-input>
                </template>
              </el-table-column>
            </el-table>
            <div class="setBtn">
              <el-button type="primary" class="BtnStyle" @click="UpdataScores">确认修改</el-button>
            </div>
          </el-main>
          <el-main class="home-footer-setting" v-if="activeMenuItem==4">
            <div class="home-footer-panel">
              <div class="home-footer-view">
                <el-table
                  :data="tableObject"
                  tooltip-effect="dark"
                  :span-method="tableObjSpanMethod"
                  border
                  stripe
                  style="width: 600px"
                >
                  <el-table-column prop="optionval" label="类目" align="center"></el-table-column>
                  <el-table-column prop="name" label="条目" header-align="center" align="center"></el-table-column>
                  <el-table-column prop="url" label="链接地址" align="center" width></el-table-column>
                  <el-table-column label="操作">
                    <template slot-scope="scope">
                      <!-- <el-button size="mini" @click="handleEdit(scope.$index, scope.row)">编辑</el-button> -->
                      <el-button
                        size="mini"
                        type="danger"
                        @click="handleDelete(scope.$index, scope.row)"
                      >删除</el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </div>
              <div class="home-footer">
                新增类目：
                <el-input v-model="itemvalue" class="new-item" placeholder="请输入内容"></el-input>
                <el-button type="danger" @click="submitAddItem">确定</el-button>
              </div>
              <div class="home-footer">
                类目选择：
                <el-select class="new-item" v-model="value" placeholder="请选择">
                  <el-option
                    v-for="item in options"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                  ></el-option>
                </el-select>条目名称：
                <el-input class="new-item" v-model="newitemname" placeholder="请输入内容"></el-input>条目链接：
                <el-input class="new-item" v-model="newitemurl" placeholder="请输入内容"></el-input>
                <el-button type="danger" @click="submitAddContent">新增条目</el-button>
              </div>
            </div>

            <!-- <div class="home-footer-panel">
              <div class="home-footer">
                更多链接信息：
                <div style="margin-top:10px" v-for="(item,index) in moreLinkInfo" :key="index">
                  链接名称：
                  <el-input class="new-item" v-model="item.title"></el-input>
                  链接地址：
                  <el-input class="new-item" v-model="item.url"></el-input>
                  <el-button @click="plusLink(index)">
                    <span style="font-size:16px">+</span>
                  </el-button>
                  <el-button @click="reduceLink(index)">
                    <span style="font-size:16px">-</span>
                  </el-button>
                  <el-button type="danger" v-if="index==moreLinkInfo.length-1">提交</el-button>
                </div>
              </div>
            </div>-->

            <div class="home-footer-panel">
              <div class="home-footer">
                备案信息：
                <el-input
                  class="new-item"
                  style="width:400px"
                  v-model="recordInfo"
                  placeholder="请输入网站备案信息"
                ></el-input>此网站版权属于：
                <el-input
                  class="new-item"
                  style="width:400px"
                  v-model="companyInfo"
                  placeholder="请输入公司名称"
                ></el-input>
                <el-button type="danger" @click="submitRecordInfo">提交</el-button>
              </div>
            </div>
          </el-main>
          <el-main class="home-footer-setting" v-if="activeMenuItem==6">
            模板名称：
            <input type="text" v-model="addTplName" />
            <el-upload
              class="upload-demo"
              action="/cms/web/add_tpl"
              :headers="headersAuthorization"
              :data="addTplData"
              name="show_pic"
            >
              <el-button size="small" type="primary">点击上传</el-button>
              <div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过1M</div>
            </el-upload>
          </el-main>
          <el-main class="home-footer-setting" v-if="activeMenuItem == 7">
            <h4>反馈列表</h4>
            <el-divider></el-divider>

            <el-table :data="tableData" style="width: 100%">
              <el-table-column prop="id" label="ID" width="180"></el-table-column>
              <el-table-column prop="describe" label="问题描述"></el-table-column>
              <el-table-column prop="contact_info" width="180" label="联系方式"></el-table-column>
              <el-table-column prop="created_at" width="180" label="创建时间"></el-table-column>
            </el-table>
          </el-main>
        </el-container>
      </el-container>
      <!-- <el-footer height="316px">
        <home-footer :logo="logoUrl"></home-footer>
      </el-footer>-->
    </el-container>
  </div>
</template>
<script>
/***
 *
 * body10    1 为组id, 2 为分类id
 * body9     综合id  可为分类id 也可为组id
 * bodt
 */
import homeFooter from "../components/homeFooter.vue"; //尾部
import IDType from "../components/IDType.vue";
import {
  getWebUser,
  getWebList,
  getUserInfo,
  getChannelList,
  getWeb,
  getModule,
  updateRowNumber,
  updataChannel,
  updateModule,
  addModuleData,
  addChannel,
  updateMoreChannel,
  signOut,
  upChannel,
  upTpl,
  getWebId,
  getBusiness,
  getBase,
  controlSwitch,
  getGroup,
  getResource,
  updateIndex,
  setIndex,
  getCategoryNode,
  setFooter,
  getFooterInfo,
  setFooterItem,
  deleteFooterItem,
  deleteFooter,
  showScore,
  updataScore,
  UniversalUpdata,
  addTpl,
  sourceMsg,
  updateKeyword,
  getSuggest
} from "../services";
export default {
  data() {
    return {
      isShow_sidebar: false,
      islogin: false,
      sidebarQQ: "",
      sidebarTel: "",
      sidebarTime: "",
      companyInfo: "",
      recordInfo: "",
      moreLinkInfo: [
        { title: "网站条款", url: "" },
        { title: "版权声明", url: "" }
      ],
      slogentext: "视觉无处不在，视觉服务中国",
      newitemname: "",
      newitemurl: "",
      options: [],
      value: "",
      tableObject: [],
      AllCategories: [],
      itemvalue: "",
      a: "2",
      ztid: "",
      ztidB: "",
      ztidC: "",
      groupid: "",
      webTitle: "", //网站title
      value3: true,
      business: [], //业务分类数据
      base: [], //基础分类数据
      itemID: null, //频道id
      userInfo: {}, //用户信息
      WebListInfo: [], //模板信息
      ChannelListInfo: [], //频道信息
      nowChannel: "", //当前频道
      nowChannelIndex: 1,
      nowChannelTitle: "", //当前频道名称
      nowLabelInfo: [], //当前标签信息
      nowModel: 1, //当前模板
      nowChannelModel: 0, // 当前频道的模板

      videoNumber: 5, //推荐视频数量，默认5个
      videoList: [], //推荐视频数据列表
      videoData: {}, //推荐视频数据
      nationalNumber: 5, // 国内数量
      overseasNumber: 5, // 海外数量
      slideshowNumber: 3, //轮播图数量
      slideNumber: 4, //未登陆轮播数量
      slideData: {}, // 未登录轮播数据
      slideList: [], // 未登录轮播列表
      showNumber: 4, //专题图数量
      showNumberB: 1, //专题策划上左大专题设置
      showNumberC: 4, //专题策划上右大专题设置
      bannerData: {}, // 轮播数据
      bannerList: [], // 轮播列表
      nationalList: [], //国内数据
      overseasList: [], // 海外数据
      hotSpecial: {}, // 热门专题数据
      hotSpecialB: {}, // 专题策划左上大专题设置
      hotSpecialC: {}, // 专题策划右上四小专题设置
      hotSpecialList: [], // 热门专题列表
      hotSpecialListB: [], // 专题测试左上大专题设置
      hotSpecialListC: [], // 专题策划右上四小专题设置
      topAdvert: {}, //顶部广告位数据
      topAdvertList: [], //顶部广告位列表
      topAdvertNumber: 0, //顶部广告位数量
      bottomAdvert: {}, //底部广告位数据
      bottomAdvertList: [], //底部广告位列表
      bottomAdvertNumber: 0, //底部广告位数量
      slideshowNumber_sure: 3, //轮播图数量---确定
      showNumber_sure: 3, //专题图数量---确定
      logoUrl: null, //logo
      logoUrlBot: null, //底部logo
      logoUrlLogin: null, //底部logo
      tplUrl: null, //模板
      tplId: 1, //要修改的模板id
      tagNumber: 1, //标签数量
      // tagNumber_sure:9,//标签数量---确定
      specialId: "", // 专题分类ID
      specialInfo: {}, // 专题频道信息
      system: {}, // 系统通知内容
      systemList: [], // 系统通知列表
      dialogVisible: false, // 通知设置对话框
      setMessage: {
        title: "",
        item: "",
        type: 0
      }, // 用于设置通知
      curSetMessage: {}, // 当前设置的通知
      moreDialogVisible: false, // 更多通知对话框
      isAddChannel: true, // 是否在新建频道,
      basicSetIsShow: false, //基本设置显示
      navColor: "#1d2b36", //导航栏底色
      navColorBot: "#282828",
      navColorSlogen: "#181A1A",
      FormData: null, //logo图片FormData对象
      logoId: "", //logoid
      logoFile: null, //logo文件

      FormDataBot: null,
      logoIdBot: "",
      logoFileBot: null,
      logoFileLogin: null,
      navId: "", //底色id
      webid: "", //租户id
      uploadCarouselIndexItem: {},
      uploadSpecialIndexItem: {},
      upType: "A",
      activeMenuItem: "",
      rowIndexAndSpan: [],
      addTplName: "",
      dynamicTags: [],
      inputVisible: false,
      inputValue: "",
      tableData: [],
      form: { ad_url: "" }
    };
  },
  created() {
    this.webid = sessionStorage.getItem("customer_id");
    console.log(this.webid);

    showScore({ web_id: this.webid }).then(res => {
      this.AllCategories = res.data;
    });
    var id = setInterval(() => {
      if (document.querySelector(".el-menu-item.is-active")) {
        clearInterval(id);
        document.querySelector(".el-menu-item.is-active").click();
      }
    }, 0);
    this.setWebId();
  },
  computed: {
    isbeijing2022() {
      return this.webid == "111130"; //冬奥组委需要隐藏
    },
    addTplData() {
      let obj = {
        web_id: this.webid,
        tpl_name: this.addTplName
      };
      return obj;
    },
    headersAuthorization() {
      const cmsToken = sessionStorage.getItem("cmsToken");
      let obj = {};
      if (cmsToken) {
        obj.Authorization = `Bearer ${cmsToken}`;
        return obj;
      }
    }
  },
  watch: {
    nowModel(val) {
      console.log("当前的啥？", val);
    },
    activeMenuItem(val) {
      if (val == 7) {
        // 获取反馈列表数据
        this.getSuggestList();
      }
    }
  },
  methods: {
    setVideoAssetID(item, index) {
      sourceMsg({ assetId: item.video_id }).then(x => {
        item["extension_info"] = x.data[0].extension_info;
        item.photo1 = item.extension_info.preview_ossid;
        item.photo2 = item.extension_info.video_image;
        this.$set(this.videoList, index, item);

        var data = {
          datas: [
            {
              id: item.id,
              module_id: 41,
              photo1: item.extension_info.preview_ossid,
              photo2: item.extension_info.video_image,
              video_id: item.video_id
            }
          ]
        };
        updateModule(data).then(res => {
          console.log(res);
        });
      });

      // this.videoList.forEach((item,index)=>{
      //   data.datas.push({
      //     id:item.id,
      //     module_id:41,
      //     video_id:item.video_id
      //   })
      // })

      //  let data = {
      //     datas: [
      //       {
      //         id: this.setMessage.item.id,
      //         module_id: this.setMessage.item.module_id,
      //         title: this.setMessage.item.title,
      //         status: "1",
      //         body: this.setMessage.item.body,
      //         resource_id: this.setMessage.item.resource_id,
      //         group_id: this.setMessage.item.group_id,
      //         category_id: this.setMessage.item.category_id,
      //         channel_id: this.nowChannel.id
      //       }
      //     ]
      //   };
      //   updateModule(data).then(res => {
      //     if (res.status_code == 1) {
      //       for (let key in this.curSetMessage) {
      //         this.curSetMessage[key] = this.setMessage.item[key];
      //       }
      //       this.$message.success("修改成功");
      //     } else {
      //       this.$message.error("修改失败，请尝试再次提交(" + res.data + ")");
      //     }
      //   });
      //设置视频资源ID
    },
    showSidebarClick(val) {
      let data = {
        web_id: this.webid,
        module_id: 23
      };
      data["body4"] = this.isShow_sidebar.toString();
      upChannel(data).then(res => {
        this.$message({
          message: res.data,
          type: "success"
        });
      });
    },
    beforeLoginClick(val) {
      let data = {
        web_id: this.webid,
        module_id: 40
      };
      data["login_index"] = val ? 1 : 0;
      upChannel(data).then(res => {
        this.$message({
          message: res.data,
          type: "success"
        });
      });
    },
    setSidebarData() {
      let data = {
        web_id: this.webid,
        module_id: 23
      };
      data["body4"] = this.isShow_sidebar.toString();
      data["body5"] = this.sidebarQQ;
      data["body6"] = this.sidebarTel;
      data["body7"] = this.sidebarTime;
      upChannel(data).then(res => {
        this.$message({
          message: res.data,
          type: "success"
        });
      });
    },
    UniversalUpdata() {
      let data = {
        module_id: 24
      };
      if (this.buy_score) {
        data["buy_score"] = this.buy_score;
      }
      if (this.login_score) {
        data["login_score"] = this.login_score;
      }
      if (this.upload_score) {
        data["upload_score"] = this.upload_score;
      }
      upChannel(data).then(res => {
        console.log(res);
        if (res.status_code == 1) {
          this.$message.success("修改成功");
        } else {
          this.$message.error("设置失败");
        }
      });
    },
    UpdataScores() {
      let data = this.AllCategories;
      var DataArr = [];
      var categoryIdArr = [];
      for (var i = 0; i < data.length; i++) {
        DataArr.push(data[i].score);
        categoryIdArr.push(data[i].id);
      }
      updataScore({
        web_id: this.webid,
        category_id: categoryIdArr,
        score: DataArr
      })
        .then(res => {
          // console.log(res);
          this.$message({
            message: "修改成功",
            type: "success"
          });
        })
        .catch(err => {
          // console.log(err);
        });
      // console.log(this.AllCategories);
      //
      // let data = {
      //   web_id: this.webid,
      //   datas: []
      // };
      //
      // this.AllCategories.forEach((item, index) => {
      //   data.datas.push({
      //     category_id: item.id,
      //     score:item.score
      //   });
      // });
      // console.log(data);
    },
    submitRecordInfo() {
      let data = {
        web_id: this.webid,
        module_id: 23
      };
      if (this.recordInfo) {
        data["body1"] = this.recordInfo;
      }
      if (this.companyInfo) {
        data["body2"] = this.companyInfo;
      }
      upChannel(data).then(res => {
        // debugger;
        this.$message({
          message: res.data,
          type: "success"
        });
      });
    },
    plusLink(index) {
      this.moreLinkInfo.splice(index + 1, 0, { title: "", url: "" });
    },
    reduceLink(index) {
      this.moreLinkInfo.splice(index, 1);
    },
    /**
    设置导航菜单位置
     */
    menuPosSet(pram) {
      let data = {
        web_id: this.webid,
        module_id: 23,
        nav: pram
      };
      upChannel(data).then(res => {
        this.$message({
          message: res.data,
          type: "success"
        });
      });
    },
    handleDelete(index, row) {
      // debugger;
      var that = this;
      deleteFooterItem({
        id: row.id,
        web_id: that.webid
      }).then(res => {
        if (res.status_code == 1) {
          // debugger;
          that.tableObject.splice(index, 1);
          var smallen = that.tableObject.filter(item => item.pid == row.pid);
          if (
            that.tableObject.filter(item => item.pid == row.pid).length == 0
          ) {
            deleteFooter({
              id: row.pid,
              web_id: that.webid
            }).then(om => {
              let index = that.options.findIndex(xx => xx.value == row.pid);
              that.options.splice(index, 1);
              that.value = "";
            });
          }
        }
      });
    },
    tableObjSpanMethod({ row, column, rowIndex, columnIndex }) {
      if (columnIndex === 0) {
        // if (rowIndex % 2 === 0) {
        //   return {
        //     rowspan: 2,
        //     colspan: 1
        //   };
        // } else {
        //   return {
        //     rowspan: 0,
        //     colspan: 0
        //   };
        // }
      }
    },
    submitAddItem() {
      var len = this.options.length;
      if (len > 5) {
        this.$message({
          message: "类目最多6个",
          type: "warning"
        });
        return;
      }
      setFooter({
        web_id: this.webid, //租户id
        service_name: this.itemvalue //名称
      }).then(res => {
        if (res.status_code == 1) {
          this.$set(this.options, len, {
            value: res.data.id,
            label: this.itemvalue
          });
          this.value = res.data.id;
          this.itemvalue = "";
        } else {
          this.$message({
            message: "新增类目失败",
            type: "warning"
          });
        }
      });
    },
    getFooter() {
      var that = this;
      getFooterInfo({ web_id: this.webid }).then(res => {
        if (res.status_code == 1) {
          res.data.forEach(item => {
            item.list.forEach(x => {
              // debugger;
              if (that.tableObject.filter(o => o.name == x.title).length == 0) {
                that.tableObject.push({
                  optionval: item.title,
                  name: x.title,
                  url: x.url,
                  id: x.id,
                  pid: item.id
                });
              }
            });
            if (that.options.filter(o => o.value == item.id).length == 0) {
              that.options.push({
                label: item.title,
                value: item.id
              });
            }
          });
        }
      });
    },
    submitAddContent() {
      setFooterItem({
        web_id: this.webid, //租户ID
        title: this.newitemname, //标题
        url: this.newitemurl, //url 需要带着http://
        service_id: this.value
      }).then(res => {
        this.getFooter();
        this.itemvalue = "";
        this.newitemname = "";
        this.newitemurl = "";
        // var len = this.tableObject.length;
        // this.$set(this.tableObject, len, {
        //   optionval: this.options.filter(x => x.value == this.value)[0].label,
        //   name: this.newitemname,
        //   url: this.newitemurl
        // });
        // this.tableObject.sort((obj1, obj2) => {
        //   var val1 = obj1.optionval;
        //   var val2 = obj2.optionval;
        //   if (val1 < val2) {
        //     return -1;
        //   } else if (val1 > val2) {
        //     return 1;
        //   } else {
        //     return 0;
        //   }
        // });
      });
    },
    /**
     * 获取轮播图尺寸
     */
    getElcarouselSize() {
      switch (this.nowModel) {
        case 44:
          return "1920x480";
        case 37:
          return "1920x480";
        case 36:
          return "897x504";
        case 29:
          return "1440x450";
        case 13:
          return "1200x950";
        case 10:
          return "1920x800";
        case 3:
          return "1600x562";
        case 1:
          return "1200x950";
        case 2:
          return "1200x950";
        default:
          return "1200x1060";
      }
    },
    uploadCarousel(item) {
      console.log(item);
      this.uploadCarouselIndexItem = item;
    },
    UploadfileElcarousel(file, fileList) {
      var data = new FormData();
      // debugger;
      if (this.uploadCarouselIndexItem.body10 == "1") {
        data.append("customSize", this.getElcarouselSize());
        data.append("groupId", this.uploadCarouselIndexItem.body9);
        data.append("asset", file.raw);
        updateIndex(data).then(async res => {
          if (res.status_code == 1) {
            let index = this.bannerList.findIndex(
              o => o.body9 == this.uploadCarouselIndexItem.body9
            );
            this.bannerList[
              index
            ].photo1 = await this.getGroupsDataByCategoryChangeV(
              this.bannerList[index]
            );
            this.$set(this.bannerList, index, this.bannerList[index]);
          }
        });
      } else if (this.uploadCarouselIndexItem.body10 == "2") {
        data.append("customSize", this.getElcarouselSize());
        data.append("category_id", this.uploadCarouselIndexItem.body9);
        data.append("asset", file.raw);
        setIndex(data).then(async res => {
          if (res.status_code == 1) {
            let index = this.bannerList.findIndex(
              o => o.body9 == this.uploadCarouselIndexItem.body9
            );
            this.bannerList[
              index
            ].photo1 = await this.getGroupsDataByCategoryChangeV(
              this.bannerList[index]
            );
            this.$set(this.bannerList, index, this.bannerList[index]);
          }
        });
      } else if (this.uploadCarouselIndexItem.body10 == "3") {
        var data = new FormData();
        let index = this.bannerList.findIndex(
          o => o.id == this.uploadCarouselIndexItem.id
        );
        data.append("promotion_cover", file.raw);
        data.append("module_id", 16);
        data.append("channel_id", this.nowChannel.id);
        data.append("id", this.uploadCarouselIndexItem.id);
        let photosrc = "";
        upChannel(data).then(res => {
          if (res.status_code == 1) {
            this.bannerList[index].photo1 = photosrc;
            this.$set(this.bannerList, index, this.bannerList[index]);
          } else {
            this.$message.error("设置失败");
          }
        });
        photosrc = window.URL.createObjectURL(file.raw);
      }
    },
    getSpecialSize(type) {
      if (type == "A") {
        switch (this.nowModel) {
          case 1:
            return "800x500";
          case 2:
            return "1200x600";
          case 3:
            return "800x500";
          case 4:
            return "1200x600";
          case 11:
            return "1200x1060";
          case 8:
            return "800x500";
          case 36:
            return "251x160";
          default:
            return "800x500";
        }
      } else if (type == "B") {
        return "900x1000";
      } else if (type == "C") {
        return "1200x800";
      }
    },
    uploadSpecial(item, type) {
      this.uploadSpecialIndexItem = item;
      this.upType = type;
    },
    UploadfileSpecial(file, fileList) {
      var data = new FormData();
      if (this.uploadSpecialIndexItem.body10 == "1") {
        data.append("customSize", this.getSpecialSize(this.upType));
        data.append("groupId", this.uploadSpecialIndexItem.body9);
        data.append("asset", file.raw);
        updateIndex(data).then(async res => {
          if (res.status_code == 1) {
            switch (this.upType) {
              case "B":
                let index = this.hotSpecialListB.findIndex(
                  o => o.body9 == this.uploadSpecialIndexItem.body9
                );
                this.hotSpecialListB[
                  index
                ].photo1 = await this.getZTDataByCategoryChangeV(
                  this.hotSpecialListB[index],
                  "B"
                );
                this.$set(
                  this.hotSpecialListB,
                  index,
                  this.hotSpecialListB[index]
                );

                break;
              case "C":
                let indexC = this.hotSpecialListC.findIndex(
                  o => o.body9 == this.uploadSpecialIndexItem.body9
                );
                this.hotSpecialListC[
                  indexC
                ].photo1 = await this.getZTDataByCategoryChangeV(
                  this.hotSpecialListC[indexC],
                  "C"
                );
                this.$set(
                  this.hotSpecialListC,
                  indexC,
                  this.hotSpecialListC[indexC]
                );
                break;
              default:
                let indexA = this.hotSpecialList.findIndex(
                  o => o.body9 == this.uploadSpecialIndexItem.body9
                );
                this.hotSpecialList[
                  indexA
                ].photo1 = await this.getZTDataByCategoryChangeV(
                  this.hotSpecialList[indexA],
                  "A"
                );
                this.$set(
                  this.hotSpecialList,
                  indexA,
                  this.hotSpecialList[indexA]
                );

                break;
            }
          }
        });
      } else if (this.uploadSpecialIndexItem.body10 == "2") {
        data.append("customSize", this.getSpecialSize(this.upType));
        data.append("category_id", this.uploadSpecialIndexItem.body9);
        data.append("asset", file.raw);
        setIndex(data).then(async res => {
          if (res.status_code == 1) {
            switch (this.upType) {
              case "B":
                let index = this.hotSpecialListB.findIndex(
                  o => o.body9 == this.uploadSpecialIndexItem.body9
                );
                this.hotSpecialListB[
                  index
                ].photo1 = await this.getZTDataByCategoryChangeV(
                  this.hotSpecialListB[index],
                  "B"
                );
                this.$set(
                  this.hotSpecialListB,
                  index,
                  this.hotSpecialListB[index]
                );

                break;
              case "C":
                let indexC = this.hotSpecialListC.findIndex(
                  o => o.body9 == this.uploadSpecialIndexItem.body9
                );
                this.hotSpecialListC[
                  indexC
                ].photo1 = await this.getZTDataByCategoryChangeV(
                  this.hotSpecialListC[indexC],
                  "C"
                );
                this.$set(
                  this.hotSpecialListC,
                  indexC,
                  this.hotSpecialListC[indexC]
                );
                break;
              default:
                let indexA = this.hotSpecialList.findIndex(
                  o => o.body9 == this.uploadSpecialIndexItem.body9
                );
                this.hotSpecialList[
                  indexA
                ].photo1 = await this.getZTDataByCategoryChangeV(
                  this.hotSpecialList[indexA],
                  "A"
                );
                this.$set(
                  this.hotSpecialList,
                  indexA,
                  this.hotSpecialList[indexA]
                );

                break;
            }
          }
        });
      } else if (this.uploadSpecialIndexItem.body10 == "3") {
        var data = new FormData();

        let moduleId = null;
        let index = null;
        switch (this.upType) {
          case "A":
            moduleId = 19;
            index = this.hotSpecialList.findIndex(
              o => o.id == this.uploadSpecialIndexItem.id
            );
            break;
          case "B":
            moduleId = 20;
            index = this.hotSpecialListB.findIndex(
              o => o.id == this.uploadSpecialIndexItem.id
            );
            break;
          case "C":
            index = this.hotSpecialListC.findIndex(
              o => o.id == this.uploadSpecialIndexItem.id
            );
            moduleId = 21;
            break;
          case "D":
            index = this.hotActivityList.findIndex(
              o => o.id == this.uploadSpecialIndexItem.id
            );
            moduleId = 26;
            break;
        }
        data.append("promotion_cover", file.raw);
        data.append("module_id", moduleId);
        data.append("channel_id", this.nowChannel.id);
        data.append("id", this.uploadSpecialIndexItem.id);
        let photosrc = "";
        upChannel(data).then(res => {
          if (res.status_code == 1) {
            switch (this.upType) {
              case "A":
                this.hotSpecialList[index].photo1 = photosrc;
                this.$set(
                  this.hotSpecialList,
                  index,
                  this.hotSpecialList[index]
                );
                break;
              case "B":
                this.hotSpecialListB[index].photo1 = photosrc;
                this.$set(
                  this.hotSpecialListB,
                  index,
                  this.hotSpecialListB[index]
                );
                break;
              case "C":
                this.hotSpecialListC[index].photo1 = photosrc;
                this.$set(
                  this.hotSpecialListC,
                  index,
                  this.hotSpecialListC[index]
                );
                break;
              case "D":
                this.hotActivityList[index].photo1 = photosrc;
                this.$set(
                  this.hotActivityList,
                  index,
                  this.hotActivityList[index]
                );
                breakl;
            }
          } else {
            this.$message.error("设置失败");
          }
        });
        photosrc = window.URL.createObjectURL(file.raw);
      }
    },
    async setWebId() {
      let data = await getWebId({
        domainName: window.location.host
      }).then();
      // console.log(data);
      this.webid = data.data;
      this.init();
    },
    //初始化信息
    init() {
      this.getWebList();
      this.getUserInfo();
      this.getChannelList();
      this.getChange();
      this.getBusiness();
      this.getBase();
      this.getWebUser();
      if (this.nowChannel.id) {
        this.getLogoUrl();
        this.getnowLabelInfo();
        this.getadvert();
        this.getSlide();
      }
    },
    //获取网站title
    getWebUser() {
      let data = {
        uid: this.webid
      };
      getWebUser(data).then(res => {
        if (res.status_code == 1) {
          this.webTitle = res.data.customName;
          this.dynamicTags = res.data.extra_config
            ? res.data.extra_config.keywords
            : [];
        }
      });
    },
    //开关状态设置
    changeSwitch(data) {
      const parse = {
        category_id: data.id + "",
        type: data.active ? 1 : 2
      };
      controlSwitch(parse).then(res => {
        if (res.status_code == 1) {
          let info = data.active ? "打开" : "关闭";
          this.$message.success(info + "成功");
        }
      });
    },
    //
    inChange() {},
    //基础设置表格列合并
    BusinessSpanMethod({ row, column, rowIndex, columnIndex }) {
      if (columnIndex === 0) {
        if (rowIndex % this.business.length === 0) {
          return {
            rowspan: this.business.length,
            colspan: 1
          };
        } else {
          return {
            rowspan: 0,
            colspan: 0
          };
        }
      }
    },
    BaseSpanMethod({ row, column, rowIndex, columnIndex }) {
      if (columnIndex === 0) {
        if (rowIndex % this.base.length === 0) {
          return {
            rowspan: this.base.length,
            colspan: 1
          };
        } else {
          return {
            rowspan: 0,
            colspan: 0
          };
        }
      }
    },

    //获取业务分类
    getBusiness() {
      let data = {
        customer_id: "111111"
      };
      getBusiness(data).then(res => {
        if (res.status_code == 1) {
          this.business = res.data.list.filter((item, index) => {
            return item.level == 1;
          });
          this.business = this.business.map(item => {
            item.active = item.active == 1 ? true : false;
            return item;
          });
        }
      });
    },
    //获取基础分类
    getBase() {
      let data = {
        customer_id: "111111"
      };
      getBase(data).then(res => {
        if (res.status_code == 1) {
          this.base = res.data.list.filter((item, index) => {
            return item.level == 1;
          });
          this.base = this.base.map(item => {
            item.active = item.active == 1 ? true : false;
            return item;
          });
        }
      });
    },
    //获取logo和底色
    getChange() {
      let data = {};
      getWeb(data).then(res => {
        const module_data = res.data.list.module_data;
        module_data.forEach(item => {
          if (item.id === 4) {
            const module_id = item.id;
            getModule({
              module_id,
              use_sort_num: 1,
              sort: "asc"
            }).then(moduleRes => {
              console.log("009-3", moduleRes);
              if (moduleRes.status_code === 1) {
                this.logoUrl = moduleRes.data.list[0].photo1;
                this.logoUrlBot = moduleRes.data.list[0].photo2;
                this.logoUrlLogin = moduleRes.data.list[0].photo15; //登录页面logo
                this.logoId = moduleRes.data.list[0].id;
              } else {
                this.$message.error(moduleRes.message);
              }
            });
          }
          if (item.id === 5) {
            const module_id = item.id;
            getModule({
              module_id,
              use_sort_num: 1,
              sort: "asc"
            }).then(moduleRes => {
              if (moduleRes.status_code === 1) {
                this.slogentext = moduleRes.data.list[0].body3;
                this.navColor = moduleRes.data.list[0].body;
                this.navColorBot = moduleRes.data.list[0].body1;
                this.navColorSlogen = moduleRes.data.list[0].body2;
                this.navId = moduleRes.data.list[0].id;
              } else {
                this.$message.error(moduleRes.message);
              }
            });
          }
          if (item.id == 23) {
            const module_id = item.id;
            getModule({
              module_id,
              use_sort_num: 1,
              sort: "asc"
            }).then(moduleRes => {
              if (moduleRes.status_code === 1) {
                // debugger;
                this.companyInfo = moduleRes.data.list[0].body2;
                this.recordInfo = moduleRes.data.list[0].body1;
                this.isShow_sidebar = eval(moduleRes.data.list[0].body4);
                this.sidebarQQ = moduleRes.data.list[0].body5;
                this.sidebarTel = moduleRes.data.list[0].body6;
                this.sidebarTime = moduleRes.data.list[0].body7;
              } else {
                this.$message.error(moduleRes.message);
              }
            });
          }
          if (item.id == 24) {
            const module_id = item.id;
            getModule({
              module_id,
              use_sort_num: 1,
              sort: "asc"
            }).then(moduleRes => {
              console.log(moduleRes);
              if (moduleRes.status_code === 1) {
                // debugger;
                this.upload_score = moduleRes.data.list[0].upload_score;
                this.login_score = moduleRes.data.list[0].login_score;
                this.buy_score = moduleRes.data.list[0].buy_score;
              } else {
                this.$message.error(moduleRes.message);
              }
            });
          }
        });
        getModule({
          module_id: 40
        }).then(res => {
          console.log("4哈", res);
          this.islogin = res.data.list[0].login_index == 1 ? true : false;
        });
      });
    },

    //获取未登录轮播
    getSlide() {
      let data = {};
      getWeb(data).then(res => {
        const module_data = res.data.list.module_data;
        module_data.forEach(item => {
          if (item.id === 15) {
            const module_id = item.id;
            getModule({
              module_id,
              use_sort_num: 1,
              sort: "asc",
              channel_id: this.nowChannel.id
            }).then(res => {
              if (res.status_code === 1) {
                this.slideData = res.data;
                this.slideList = res.data.list;
                this.slideNumber = res.data.list.length;
              } else {
                this.$message.error(res.message);
              }
            });
          }
        });
      });
    },
    //获取广告位
    getadvert() {
      let data = {};
      getWeb(data).then(res => {
        const module_data = res.data.list.module_data;
        module_data.forEach(item => {
          if (item.id === 11) {
            const module_id = item.id;
            getModule({
              module_id,
              use_sort_num: 1,
              sort: "asc",
              channel_id: this.nowChannel.id
            }).then(res => {
              if (res.status_code === 1) {
                this.topAdvert = res.data;
                this.topAdvertList = res.data.list;
                this.topAdvertNumber = res.data.list.length;
              } else {
                this.$message.error(res.message);
              }
            });
          }
          if (item.id === 12) {
            const module_id = item.id;
            getModule({
              module_id,
              use_sort_num: 1,
              sort: "asc",
              channel_id: this.nowChannel.id
            }).then(res => {
              if (res.status_code === 1) {
                this.bottomAdvert = res.data;
                this.bottomAdvertList = res.data.list;
                this.bottomAdvertNumber = res.data.list.length;
              } else {
                this.$message.error(res.message);
              }
            });
          }
        });
      });
    },
    //退出登陆
    signOut() {
      signOut().then(res => {
        sessionStorage.removeItem("cmsUser");
        sessionStorage.removeItem("cmsToken");
        if (this.$route.fullPath == "/") {
          location.reload();
          this.$message.success("退出成功");
          location.href = location.href + "login";
        } else {
          this.$router.push({ path: "/" });
        }
      });
    },

    // **************************************************************************************
    set111() {
      this.basicSetIsShow = true;
    },
    uploadLogoLogin(event) {
      var target = event.srcElement ? event.srcElement : event.target;
      if (target.files[0]) {
        var file = target.files[0];
        var sr = window.URL.createObjectURL(file);
        this.logoUrlLogin = sr;
        var data = new FormData();
        data.append("photo15", file);
        this.FormDataBot = data;
        this.logoFileLogin = file;
      }
    },
    uploadLogoBot(event) {
      var target = event.srcElement ? event.srcElement : event.target;
      console.log(target.files);
      if (target.files[0]) {
        var file = target.files[0];
        var sr = window.URL.createObjectURL(file);
        this.logoUrlBot = sr;
        var data = new FormData();
        data.append("photo2", file);
        this.FormDataBot = data;
        this.logoFileBot = file;
      }
    },

    uploadSpread(event, item, id) {
      var target = event.srcElement ? event.srcElement : event.target;
      if (target.files[0]) {
        var file = target.files[0];
        var data = new FormData();

        data.append("promotion_pic", file);
        data.append("module_id", id);
        data.append("channel_id", this.nowChannel.id);
        data.append("id", item.id);
        upChannel(data).then(res => {
          if (res.status_code == 1) {
            this.$message.success("修改成功");
          } else {
            this.$message.error("logo设置失败");
          }
        });
      }
    },
    //logo上传
    uploadLogo(event) {
      var target = event.srcElement ? event.srcElement : event.target;
      if (target.files[0]) {
        var file = target.files[0];
        var sr = window.URL.createObjectURL(file);
        this.logoUrl = sr;
        var data = new FormData();
        data.append("photo1", file);
        this.FormData = data;
        this.logoFile = file;
      }
    },
    //顶部广告上传
    upTopAdvert($event, item) {
      var target = $event.srcElement ? $event.srcElement : $event.target;
      if (target.files[0]) {
        var file = target.files[0];
        var sr = window.URL.createObjectURL(file);
        item.link_url = sr;
        var data = new FormData();
        data.append("link_url", file);
        this.FormData = data;
        item.file = file;
        var data1 = new FormData();
        data1.append("id", item.id);
        data1.append("module_id", 11);
        data1.append("photo1", item.file);
        data1.append("body", 1);
        upChannel(data1).then(res => {
          if (res.status_code == 1) {
            this.$message.success("上传成功");
          } else {
            this.$message.error("上传失败");
          }
        });
      }
    },
    //底部广告上传
    upBottomAdvert($event, item) {
      var target = $event.srcElement ? $event.srcElement : $event.target;
      if (target.files[0]) {
        var file = target.files[0];
        var sr = window.URL.createObjectURL(file);
        item.link_url = sr;
        var data = new FormData();
        data.append("link_url", file);
        this.FormData = data;
        item.file = file;
        var data1 = new FormData();
        data1.append("id", item.id);
        data1.append("module_id", 12);
        data1.append("photo1", item.file);
        data1.append("body", 1);
        upChannel(data1).then(res => {
          if (res.status_code == 1) {
            this.$message.success("上传成功");
          } else {
            this.$message.error("上传失败");
          }
        });
      }
    },
    // 频道排序
    sortChannel(type, index) {
      let oldData = JSON.parse(JSON.stringify(this.ChannelListInfo));
      if (type == "up") {
        [this.ChannelListInfo[index], this.ChannelListInfo[index - 1]] = [
          this.ChannelListInfo[index - 1],
          this.ChannelListInfo[index]
        ];
      } else if (type == "down") {
        [this.ChannelListInfo[index], this.ChannelListInfo[index + 1]] = [
          this.ChannelListInfo[index + 1],
          this.ChannelListInfo[index]
        ];
      }
      this.ChannelListInfo = [...this.ChannelListInfo];
      this.ChannelListInfo.forEach((item, index) => {
        item.sort = index;
        item.channel_id = item.id;
        item.channel_name = item.page_name;
      });
      updateMoreChannel({ datas: this.ChannelListInfo }).then(res => {
        if (res.status_code == 1) {
          this.$message.success("修改成功");
        } else {
          this.ChannelListInfo = [...oldData];
          this.$message.error("修改失败，请尝试重新提交");
        }
      });
    },
    // 删除频道
    deleteChannel(item, index) {
      let data = {
        channel_id: item.id,
        tpl_id: item.tpl_id,
        status: "2"
      };
      updataChannel(data).then(res => {
        if (res.status_code == 1) {
          this.ChannelListInfo.splice(index, 1);
          this.$message.success("删除成功");
          document.querySelector(".el-menu-item").click();
        } else {
          this.$message.error("删除失败，请重新尝试");
        }
      });
    },
    // 删除通知
    deleteMessage(item, index) {
      let data = {
        datas: [
          {
            id: item.id,
            module_id: item.module_id,
            status: "2",
            channel_id: this.nowChannel.id
          }
        ]
      };
      updateModule(data).then(res => {
        if (res.status_code == 1) {
          this.systemList.splice(index, 1);
          this.$message.success("修改成功");
        } else {
          this.$message.error("修改失败，请尝试再次提交(" + res.data + ")");
        }
      });
    },
    // 提交通知的设置
    submeMessage() {
      this.dialogVisible = false;
      if (this.setMessage.type) {
        // 新建
        let data = {
          module_id: 7,
          title: this.setMessage.item.title,
          body: this.setMessage.item.body,
          channel_id: this.nowChannel.id
        };
        addModuleData(data).then(res => {
          if (res.status_code == 1) {
            getModule({
              module_id: data.module_id,
              sort: "desc",
              channel_id: this.nowChannel.id
            }).then(res => {
              if (res.status_code == 1) {
                this.system = res.data;
                this.systemList = res.data.list;
                this.$message.success("修改成功");
              }
            });
          } else {
            this.$message.error("修改失败，请尝试再次提交");
          }
        });
      } else {
        // 编辑
        let data = {
          datas: [
            {
              id: this.setMessage.item.id,
              module_id: this.setMessage.item.module_id,
              title: this.setMessage.item.title,
              status: "1",
              body: this.setMessage.item.body,
              resource_id: this.setMessage.item.resource_id,
              group_id: this.setMessage.item.group_id,
              category_id: this.setMessage.item.category_id,
              channel_id: this.nowChannel.id
            }
          ]
        };
        updateModule(data).then(res => {
          if (res.status_code == 1) {
            for (let key in this.curSetMessage) {
              this.curSetMessage[key] = this.setMessage.item[key];
            }
            this.$message.success("修改成功");
          } else {
            this.$message.error("修改失败，请尝试再次提交(" + res.data + ")");
          }
        });
      }
    },
    // 通知设置弹框
    openMe({ type = 0, item = { title: "", body: "" } }) {
      this.setMessage.title = type ? "新建通知" : "编辑通知";
      this.setMessage.item = JSON.parse(JSON.stringify(item));
      this.curSetMessage = item;
      this.setMessage.type = type;
      this.dialogVisible = true;
    },
    // 设置专题分类ID
    setShowType() {
      this.init();
      var data = {
        datas: [
          {
            id: this.specialInfo.id,
            module_id: 8,
            category_id: this.specialId,
            channel_id: this.nowChannel.id
          }
        ]
      };
      updateModule(data).then(res => {
        if (res.status_code == 1) {
          this.$message.success("修改成功");
        } else {
          this.$message.error("修改失败，请尝试再次提交(" + res.data + ")");
        }
      });
    },
    // 控制显示模板
    moduleShow(index, item) {
      let show = true;
      if (item.id == 1) {
        show = this.nowChannelIndex == 0 ? true : false;
      } else if (item.id == 4 || item.id == 8) {
        show = this.nowChannelIndex == 0 ? false : true;
      } else {
        show = true;
      }
      return show;
    },
    // 新增频道
    addChannel() {
      if (this.nowChannelTitle == "首页") {
        this.nowChannelIndex = 0;
      }
      this.basicSetIsShow = false;
      this.nowChannel = {};
      this.nowChannelTitle = "";
      this.nowModel = 1;
      this.nowChannelIndex = 0;
      this.isAddChannel = true;
    },
    // 设置当前在编辑的频道
    setCurPage(item, index) {
      console.log("看下当前频道数据", item);
      this.basicSetIsShow = false;
      this.nowChannel = item;
      this.nowChannelTitle = item.page_name;
      this.nowChannelModel = item.tpl_id;
      this.nowModel = item.tpl_id;
      this.nowChannelIndex = index;
      this.isAddChannel = false;
      this.getLogoUrl();
      this.getnowLabelInfo();
      this.getChange();
      this.getadvert();
      this.getSlide();
    },
    //获取当前用户信息
    getUserInfo() {
      getUserInfo().then(res => {
        this.userInfo = res.data;
      });
    },
    //获取模板信息
    getWebList() {
      let data = {};
      getWebList(data).then(res => {
        if (res.status_code === 1) {
          this.WebListInfo = res.data;
        } else {
          this.$message.error(res.message);
        }
      });
    },
    //获取网站的频道列表
    getChannelList() {
      getChannelList({ user_sort: 1 }).then(res => {
        this.ChannelListInfo = res.data;
        if (this.isAddChannel) {
          let item = this.ChannelListInfo.find(item => {
            if (item.page_name == this.nowChannelTitle) return item;
          });
          if (item) {
            let index = this.ChannelListInfo.indexOf(item);
            document.querySelector(".el-submenu__title").click();
            var id = setInterval(() => {
              if (document.querySelectorAll(".el-menu-item")[index]) {
                clearInterval(id);
                document.querySelectorAll(".el-menu-item")[index].click();
              }
            }, 0);
          }
        }
      });
    },
    //获取标签内容
    getnowLabelInfo() {
      let data = {};
      getWeb(data).then(res => {
        const module_data = res.data.list.module_data;
        module_data.forEach(item => {
          if (item.title === "标签") {
            const module_id = item.id;
            getModule({
              module_id,
              use_sort_num: 1,
              channel_id: this.nowChannel.id,
              sort: "asc"
            }).then(moduleRes => {
              if (moduleRes.status_code === 1) {
                this.nowLabelInfo = moduleRes.data.list;
                this.tagNumber = moduleRes.data.list.length;
              } else {
                this.$message.error(moduleRes.message);
              }
            });
          }
        });
      });
    },
    //获取logoURL
    getLogoUrl() {
      let data = {};
      getWeb(data).then(res => {
        const module_data = res.data.list.module_data;
        console.log("009-1", res);
        module_data.forEach(item => {
          if (item.title === "logo") {
            const module_id = 1;
            getModule({
              module_id,
              use_sort_num: 1,
              channel_id: this.nowChannel.id
            }).then(moduleRes => {
              if (moduleRes.status_code === 1) {
                console.log("009-2", moduleRes);
                this.logoUrl = moduleRes.data.list[0].photo1;
                this.logoUrlBot = moduleRes.data.list[0].photo2;
                this.logoUrlLogin = moduleRes.data.list[0].photo15;
              } else {
                this.$message.error(moduleRes.message);
              }
            });
          } else if (item.title === "A轮播图") {
            let module_id = item.id;
            getModule({
              module_id,
              use_sort_num: 1,
              channel_id: this.nowChannel.id
            }).then(async res => {
              if (res.status_code == 1) {
                this.bannerData = res.data;
                this.bannerList = res.data.list;
                this.slideshowNumber = this.bannerList.length;
                for (let i = 0; i < this.bannerList.length; i++) {
                  this.bannerList[
                    i
                  ].photo1 = await this.getGroupsDataByCategoryChangeV(
                    this.bannerList[i]
                  );
                  this.$set(this.bannerList, i, this.bannerList[i]);
                }
              }
            });
          } else if (item.title === "B轮播图") {
            // 用作国内推荐
            let module_id = item.id;
            getModule({
              module_id,
              use_sort_num: 1,
              channel_id: this.nowChannel.id
            }).then(async res => {
              if (res.status_code == 1) {
                this.nationalList = res.data.list;
                this.nationalNumber = this.nationalList.length;
                for (let i = 0; i < this.nationalList.length; i++) {
                  this.nationalList[
                    i
                  ].photo1 = await this.getGroupsDataByCategoryChangeV(
                    this.nationalList[i]
                  );
                  this.$set(this.nationalList, i, this.nationalList[i]);
                }
              }
            });
          } else if (item.title === "C轮播图") {
            // 用作国际推荐
            let module_id = item.id;
            getModule({
              module_id,
              use_sort_num: 1,
              channel_id: this.nowChannel.id
            }).then(async res => {
              if (res.status_code == 1) {
                this.overseasList = res.data.list;
                this.overseasNumber = this.overseasList.length;
                for (let i = 0; i < this.overseasList.length; i++) {
                  this.overseasList[
                    i
                  ].photo1 = await this.getGroupsDataByCategoryChangeV(
                    this.overseasList[i]
                  );
                  this.$set(this.overseasList, i, this.overseasList[i]);
                }
              }
            });
          } else if (item.title === "A热门专题") {
            let module_id = 19;
            getModule({
              module_id,
              use_sort_num: 1,
              channel_id: this.nowChannel.id
            }).then(async res => {
              if (res.status_code === 1) {
                this.hotSpecial = res.data;
                this.hotSpecialList = res.data.list;
                this.showNumber = this.hotSpecialList.length;
                for (let i = 0; i < this.hotSpecialList.length; i++) {
                  this.hotSpecialList[
                    i
                  ].photo1 = await this.getZTDataByCategoryChangeV(
                    this.hotSpecialList[i],
                    "A"
                  );
                  this.$set(this.hotSpecialList, i, this.hotSpecialList[i]);
                }
              }
            });
          } else if (item.title == "B热门专题") {
            //  专题策划上左大专题设置
            let module_id = 20;
            getModule({
              module_id,
              use_sort_num: 1,
              channel_id: this.nowChannel.id
            }).then(async res => {
              if (res.status_code === 1) {
                this.hotSpecialB = res.data;
                this.hotSpecialListB = res.data.list.slice(0, 1);
                this.showNumberB = 1;
                if (this.hotSpecialListB.length > 0) {
                  for (let i = 0; i < 1; i++) {
                    this.hotSpecialListB[
                      i
                    ].photo1 = await this.getZTDataByCategoryChangeV(
                      this.hotSpecialListB[i],
                      "B"
                    );
                    this.$set(this.hotSpecialListB, i, this.hotSpecialListB[i]);
                  }
                }
              }
            });
          } else if (item.title == "C热门专题") {
            //专题策划上右四个小专题设置
            let module_id = 21;
            getModule({
              module_id,
              use_sort_num: 1,
              channel_id: this.nowChannel.id
            }).then(async res => {
              if (res.status_code === 1) {
                this.hotSpecialC = res.data;
                this.hotSpecialListC = res.data.list;
                this.showNumberC = this.hotSpecialListC.length;
                for (let i = 0; i < this.hotSpecialListC.length; i++) {
                  this.hotSpecialListC[
                    i
                  ].photo1 = await this.getZTDataByCategoryChangeV(
                    this.hotSpecialListC[i],
                    "C"
                  );
                  this.$set(this.hotSpecialListC, i, this.hotSpecialListC[i]);
                }
              }
            });
          } else if (item.title === "专题频道") {
            let module_id = 8;
            getModule({
              module_id,
              use_sort_num: 1,
              channel_id: this.nowChannel.id
            }).then(res => {
              if (res.status_code === 1) {
                this.specialInfo = res.data.list[0];
                this.specialId = this.specialInfo
                  ? this.specialInfo.category_id
                  : "";
              }
            });
          } else if (item.title == "推荐视频") {
            let module_id = 41;
            getModule({
              module_id,
              use_sort_num: 1,
              channel_id: this.nowChannel.id
            }).then(res => {
              if (res.status_code === 1) {
                this.videoData = res.data;
                this.videoList = res.data.list;
                this.videoNumber = res.data.list.length;
              }
            });
          } else if (item.title === "系统通知") {
            let module_id = 7;
            getModule({
              module_id,
              sort: "desc",
              channel_id: this.nowChannel.id
            }).then(res => {
              if (res.status_code == 1) {
                this.system = res.data;
                this.systemList = res.data.list;
              }
            });
          }
        });
      });
    },
    //设置标签数量tagNumber
    setTagNumber() {
      const data = {
        module_id: 1,
        row_num: this.tagNumber,
        channel_id: this.nowChannel.id
      };
      updateRowNumber(data).then(res => {
        if (res.status_code === 1) {
          this.$message.success("修改成功");
          this.nowLabelInfo = res.data.list;
        } else {
          this.$message.error(res.message);
        }
      });
    },
    // 设置模板缩略图
    setModulImg(id, index, $event) {
      if ($event.target.files[0]) {
        let file = $event.target.files[0];
        let url = URL.createObjectURL(file);
        let data = new FormData();
        data.append("web_id", this.webid);
        data.append("tpl_id", id);
        data.append("show_pic", file);
        upTpl(data).then(res => {
          if (res.status_code == 1) {
            document.querySelector(
              `.modelChoose li:nth-of-type(${index + 1}) img`
            ).src = url;
            this.$message.success("修改成功");
          } else {
            this.$message.error("修改失败，请尝试再次提交");
          }
        });
      }
    },
    //标签排序设置
    sort_down(index) {
      //下
      const arr = this.nowLabelInfo;
      const item1 = arr[index];
      const item2 = arr[index + 1];
      arr[index + 1] = item1;
      arr[index] = item2;
      this.nowLabelInfo = [...arr];
    },
    sort_up(index) {
      //上
      const arr = this.nowLabelInfo;
      const item1 = arr[index];
      const item2 = arr[index - 1];
      arr[index - 1] = item1;
      arr[index] = item2;
      this.nowLabelInfo = [...arr];
    },
    // 确定修改频道名称
    setChannerName(type) {
      if (this.isAddChannel) {
        let data = {
          tpl_id: 1,
          channel_name: this.nowChannelTitle,
          channel_id: this.nowChannel.id
        };
        addChannel(data).then(res => {
          if (res.status_code == 1) {
            this.$message.success("添加频道成功");
            this.getChannelList();
          } else {
            this.$message.error("添加失败，请尝试再次提交");
          }
        });
      } else {
        let name;
        if (type) {
          let channel = this.ChannelListInfo.find(item => {
            if (item.id == this.nowChannel.id) return item;
          });
          name = channel.page_name;
        } else {
          name = this.nowChannelTitle;
        }
        let data = {
          channel_id: this.nowChannel.id,
          channel_name: name,
          tpl_id: this.nowModel,
          sort: this.nowChannel.sort
        };
        updataChannel(data).then(res => {
          if (res.status_code == 1) {
            let channel = this.ChannelListInfo.find(item => {
              if (item.id == this.nowChannel.id) return item;
            });
            channel.page_name = this.nowChannelTitle;
            channel.tpl_id = this.nowModel;
            this.nowChannel.tpl_id = this.nowModel;
            this.$message.success("修改成功");
            //  location.reload();
            this.init();
          } else {
            this.$message.error("修改失败，请尝试再次提交");
          }
        });
      }
    },
    handelSelect(index) {
      // debugger;
      this.activeMenuItem = index;
      const reg = /^1-/;
      if (index.match(reg)) {
        const nowChannel = parseInt(index.charAt(index.length - 1));
        this.nowChannel = nowChannel;
      }
      if (index == 4) {
        this.getFooter();
      }
    },

    //轮播图系列
    async getGroupsDataByCategoryChangeV(item) {
      var obj = {
        review_state: 1 + "",
        search_type: 2 + "",
        // online_state: 1 + "",
        customer_id: this.webid,
        group_id: item.body9
      };
      if (item.body10 == "1") {
        let resul = await getResource(obj, 2);
        if (resul.data) {
          return this.getCarouselShowImgUrl(resul.data.data[0]);
        } else {
          return "";
        }
      } else if (item.body10 == "2") {
        let resul = await getCategoryNode({
          customer_id: this.webid,
          category_id: item.body9
        });
        if (resul.data) {
          return this.getCarouselShowImgUrlCategory(resul.data[0]);
        } else {
          return "";
        }
      } else if (item.body10 == "3") {
        return item.promotion_cover;
      }
    },
    getCarouselShowImgUrlCategory(item) {
      if (item.custom_index_image) {
        if (item.custom_index_image[this.getElcarouselSize()]) {
          return item.custom_index_image[this.getElcarouselSize()];
        }
      }
      return item.index_image;
    },
    getCarouselShowImgUrl(item) {
      if (item.group_custom_index_image) {
        if (item.group_custom_index_image[this.getElcarouselSize()]) {
          return item.group_custom_index_image[this.getElcarouselSize()];
        }
      }
      return item.group_index_image;
    },

    //专题系列
    async getZTDataByCategoryChangeV(item, type) {
      if (item.body10 == "1") {
        var obj = {
          review_state: 1 + "",
          search_type: 2 + "",
          // online_state: 1 + "",
          customer_id: this.webid,
          group_id: item.body9
        };
        let resul = await getResource(obj, 2);
        if (resul.data) {
          return this.getSpecialShowImgUrl(resul.data.data[0], type);
        } else {
          return "";
        }
      } else if (item.body10 == "2") {
        let resul = await getCategoryNode({
          customer_id: this.webid,
          category_id: item.body9
        });
        if (resul.data) {
          return this.getSpecialShowImgUrlCategory(resul.data[0], type);
        } else {
          return "";
        }
      } else if (item.body10 == "3") {
        return item.promotion_cover;
      }
    },
    getSpecialShowImgUrlCategory(item, type) {
      if (item.custom_index_image) {
        if (item.custom_index_image[this.getSpecialSize(type)]) {
          return item.custom_index_image[this.getSpecialSize(type)];
        }
      }
      return item.index_image;
    },
    getSpecialShowImgUrl(item, type) {
      if (item.group_custom_index_image) {
        if (item.group_custom_index_image[this.getSpecialSize(type)]) {
          return item.group_custom_index_image[this.getSpecialSize(type)];
        }
      }
      return item.group_index_image;
    },

    //轮播设置
    setSlideshowNumber() {
      let data = {
        module_id: 16,
        row_num: this.slideshowNumber,
        channel_id: this.nowChannel.id
      };
      updateRowNumber(data).then(res => {
        if (res.status_code == 1) {
          this.bannerData = res.data;
          this.bannerList = res.data.list;
          this.slideshowNumber = res.data.list.length;
          this.$message.success("修改成功");
        } else {
          this.$message.error("修改失败，请尝试重新提交");
        }
      });
    },
    // 国内设置 B轮播图
    setNationalNumber() {
      let data = {
        module_id: 17,
        row_num: this.nationalNumber,
        channel_id: this.nowChannel.id
      };
      updateRowNumber(data).then(res => {
        if (res.status_code == 1) {
          this.nationalList = res.data.list;
          this.nationalNumber = res.data.list.length;
          this.$message.success("修改成功");
        } else {
          this.$message.error("修改失败，请尝试重新提交");
        }
      });
    },
    // 国际设置 C轮播图
    setOverseasNumber() {
      let data = {
        module_id: 18,
        row_num: this.overseasNumber,
        channel_id: this.nowChannel.id
      };
      updateRowNumber(data).then(res => {
        if (res.status_code == 1) {
          this.overseasList = res.data.list;
          this.overseasNumber = res.data.list.length;
          this.$message.success("修改成功");
        } else {
          this.$message.error("修改失败，请尝试重新提交");
        }
      });
    },

    setVideoNumber() {
      //设置视频个数
      if (this.videoNumber !== this.videoList.length) {
        let data = {
          module_id: 41,
          row_num: this.videoNumber,
          channel_id: this.nowChannel.id
        };
        updateRowNumber(data).then(res => {
          if (res.status_code == 1) {
            this.videoData = res.data;
            this.videoList = res.data.list;
            this.videoNumber = res.data.list.length;
            this.$message.success("修改成功");
          } else {
            this.$message.error("修改失败，请尝试重新提交");
          }
        });
      }
    },
    //未登录轮播图设置
    setSlideNumber() {
      if (this.slideNumber !== this.slideList.length) {
        let data = {
          module_id: 15,
          row_num: this.slideNumber,
          channel_id: this.nowChannel.id
        };

        updateRowNumber(data).then(res => {
          if (res.status_code == 1) {
            this.slideData = res.data;
            this.slideList = res.data.list;
            this.slideNumber = res.data.list.length;
            this.$message.success("修改成功");
          } else {
            this.$message.error("修改失败，请尝试重新提交");
          }
        });
      }
    },
    //专题设置
    setShowNumber() {
      // if (this.showNumber !== this.hotSpecialList.length) {
      let data = {
        module_id: 19,
        row_num: this.showNumber,
        channel_id: this.nowChannel.id
      };
      updateRowNumber(data).then(res => {
        if (res.status_code == 1) {
          this.hotSpecial = res.data;
          this.hotSpecialList = res.data.list;
          this.showNumber = res.data.list.length;
          this.$message.success("修改成功");
        } else {
          this.$message.error("修改失败，请尝试重新提交");
        }
      });
    },
    //专题策划左上大专题设置
    setShowNumberB() {
      let data = {
        module_id: 20,
        row_num: this.showNumberB,
        channel_id: this.nowChannel.id
      };

      updateRowNumber(data).then(res => {
        if (res.status_code == 1) {
          this.hotSpecialB = res.data;
          this.hotSpecialListB = res.data.list;
          this.showNumberB = res.data.list.length;
          this.$message.success("修改成功");
        } else {
          this.$message.error("修改失败，请尝试重新提交");
        }
      });
    },
    //专题策划右上四小专题设置
    setShowNumberC() {
      // if (this.showNumberC !== this.hotSpecialListC.length) {
      let data = {
        module_id: 21,
        row_num: this.showNumberC,
        channel_id: this.nowChannel.id
      };

      updateRowNumber(data).then(res => {
        if (res.status_code == 1) {
          this.hotSpecialC = res.data;
          this.hotSpecialListC = res.data.list;
          this.showNumberC = res.data.list.length;
          this.$message.success("修改成功");
        } else {
          this.$message.error("修改失败，请尝试重新提交");
        }
      });
    },
    //顶部广告位设置
    setTopAdvertNumber() {
      if (this.topAdvertNumber !== this.topAdvertList.length) {
        let data = {
          module_id: 11,
          row_num: this.topAdvertNumber,
          channel_id: this.nowChannel.id
        };
        updateRowNumber(data).then(res => {
          if (res.status_code == 1) {
            this.topAdvert = res.data;
            this.topAdvertList = res.data.list;
            this.topAdvertNumber = res.data.list.length;
            this.$message.success("修改成功");
          } else {
            this.$message.error("修改失败，请尝试重新提交");
          }
        });
      }
    },
    //底部广告位设置
    setBottomAdvertNumber() {
      if (this.bottomAdvertNumber !== this.bottomAdvertList.length) {
        let data = {
          module_id: 12,
          row_num: this.bottomAdvertNumber,
          channel_id: this.nowChannel.id
        };
        updateRowNumber(data).then(res => {
          if (res.status_code == 1) {
            this.bottomAdvert = res.data;
            this.bottomAdvertList = res.data.list;
            this.bottomAdvertNumber = res.data.list.length;
            this.$message.success("修改成功");
          } else {
            this.$message.error("修改失败，请尝试重新提交");
          }
        });
      }
    },
    modelChoose(type) {
      this.nowModel = type;
      this.setChannerName(1);
    },
    setOrder_up() {},
    //提交所有设置
    setAll() {
      var data = {
        datas: []
      };

      this.videoList.forEach((item, index) => {
        data.datas.push({
          id: item.id,
          module_id: 41,
          video_id: item.video_id
        });
      });
      this.nowLabelInfo.forEach((item, index) => {
        data.datas.push({
          id: item.id,
          module_id: 1,
          title: item.title,
          sort_num: index + 1 + "",
          status: "1",
          resource_id: item.resource_id,
          group_id: item.group_id,
          category_id: item.category_id
        });
      });
      this.bannerList.forEach((item, index) => {
        data.datas.push({
          id: item.id,
          module_id: 16,
          title: item.title,
          link_url: item.link_url,
          status: "1",
          resource_id: item.resource_id,
          group_id: item.group_id,
          category_id: item.category_id,
          body10: item.body10,
          body9: item.body9
        });
      });
      // debugger;
      this.slideList.forEach((item, index) => {
        data.datas.push({
          id: item.id,
          module_id: 15,
          title: item.title,
          link_url: item.link_url,
          status: "1",
          resource_id: item.resource_id,
          group_id: item.group_id,
          category_id: item.category_id
        });
      });
      this.hotSpecialListB.forEach((item, index) => {
        data.datas.push({
          id: item.id,
          module_id: 20,
          title: item.title,
          status: "1",
          resource_id: item.resource_id,
          group_id: item.group_id,
          category_id: item.category_id,
          body10: item.body10,
          body9: item.body9
        });
      });
      this.hotSpecialListC.forEach((item, index) => {
        data.datas.push({
          id: item.id,
          module_id: 21,
          title: item.title,
          status: "1",
          resource_id: item.resource_id,
          group_id: item.group_id,
          category_id: item.category_id,
          body10: item.body10,
          body9: item.body9
        });
      });
      this.hotSpecialList.forEach((item, index) => {
        data.datas.push({
          id: item.id,
          module_id: 19,
          title: item.title,
          status: "1",
          resource_id: item.resource_id,
          group_id: item.group_id,
          category_id: item.category_id,
          body10: item.body10,
          body9: item.body9
        });
      });
      this.topAdvertList.forEach((item, index) => {
        data.datas.push({
          id: item.id,
          module_id: 11,
          title: item.title,
          status: "1",
          link_url: item.link_url,
          resource_id: item.resource_id,
          group_id: item.group_id,
          category_id: item.category_id
        });
      });
      this.bottomAdvertList.forEach((item, index) => {
        data.datas.push({
          id: item.id,
          module_id: 12,
          title: item.title,
          status: "1",
          link_url: item.link_url,
          resource_id: item.resource_id,
          group_id: item.group_id,
          category_id: item.category_id
        });
      });
      updateModule(data).then(res => {
        if (res.status_code == 1) {
          this.$message.success("修改成功");
          this.init();
        } else {
          this.$message.error("修改失败，请尝试再次提交(" + res.data + ")");
        }
      });
    },
    //取消所有设置
    removeAll() {
      this.init();
    },
    //提交模板图片
    setTpl() {
      var data = new FormData();
      data.append("tpl_id", this.tplId);
      data.append("show_pic", this.tplFile);
      if (this.tplFile !== null) {
        upTpl(data).then(res => {
          if (res.status_code == 1) {
            this.$message.success("修改成功!");
            this.getWebList();
          } else {
            this.$message.error("模板设置失败!");
          }
        });
      } else {
        this.$message.error("没有上传模板图片");
      }
    },
    //基本设置提交logo
    setLogo() {
      var data1 = new FormData();
      data1.append("id", this.logoId);
      data1.append("module_id", 4);
      if (this.logoFile) {
        data1.append("photo1", this.logoFile);
      }
      if (this.logoFileBot) {
        data1.append("photo2", this.logoFileBot);
      }
      if (this.logoFileLogin) {
        data1.append("photo15", this.logoFileLogin);
      }
      data1.append("body", 1);
      if (
        this.logoFile !== null ||
        this.logoUrlBot != null ||
        this.logoFileLogin != null
      ) {
        upChannel(data1).then(res => {
          if (res.status_code == 1) {
            this.$message.success("修改成功");
          } else {
            this.$message.error("logo设置失败");
          }
        });
      } else {
        this.$message.error("没有上传图片logo");
      }
    },
    setslogentext() {
      let data = {
        id: this.navId,
        module_id: 5
      };
      if (this.slogentext) {
        data["body3"] = this.slogentext;
      }
      upChannel(data).then(res => {
        if (res.status_code == 1) {
          this.$message.success("设置成功");
        } else {
          this.$message.error("设置失败");
        }
      });
    },
    //基本设置提交首页底色
    setColor() {
      let data = {
        id: this.navId,
        module_id: 5
      };
      if (this.navColor) {
        data["body"] = this.navColor;
      }
      if (this.navColorBot) {
        data["body1"] = this.navColorBot;
      }
      if (this.navColorSlogen) {
        data["body2"] = this.navColorSlogen;
      }
      upChannel(data).then(res => {
        if (res.status_code == 1) {
          this.$message.success("颜色修改成功");
        } else {
          this.$message.error("颜色设置失败");
        }
      });
    },
    //取消logo设置
    removelogo() {
      var data = new FormData();
      data.append("id", this.logoId);
      data.append("module_id", 4);
      data.append("body", 2);
      upChannel(data).then(res => {
        if (res.status_code == 1) {
          this.$message.success("取消logo成功");
        } else {
          this.$message.error("取消logo失败");
        }
      });
    },
    handleClose(tag) {
      this.dynamicTags.splice(this.dynamicTags.indexOf(tag), 1);
    },

    showInput() {
      this.inputVisible = true;
      this.$nextTick(_ => {
        this.$refs.saveTagInput.$refs.input.focus();
      });
    },

    handleInputConfirm() {
      let inputValue = this.inputValue;
      if (inputValue) {
        this.dynamicTags.push(inputValue);
      }
      this.inputVisible = false;
      this.inputValue = "";
    },
    handleSaveKeywords() {
      const data = {
        data: this.dynamicTags
      };
      updateKeyword(data).then(async res => {
        if (res.status_code == 1) {
          this.$message.success("保存关键词成功");
        } else {
          this.$message.error("保存关键词失败，请重试");
        }
      });
    },
    async getSuggestList() {
      const params = { web_id: this.webid };
      await getSuggest(params).then(res => {
        if (res.status_code == 1) {
          this.tableData = res.data;
        } else {
          this.$message.error("获取数据失败，请稍后再试");
        }
      });
    },
    handleBlurUrl(item) {
      const data = {
        ad_url: this.form.ad_url,
        id: item.id,
        module_id: item.module_id
      };
      upChannel(data).then(res => {
        if (res.status_code == 1) {
          this.$message.success("保存成功");
        } else {
          this.$message.error("保存失败");
        }
      });
    }
  },
  components: {
    homeFooter,
    IDType
  }
};
</script>

<style lang="less">
.home-footer-setting {
  .home-footer-panel {
    margin-top: 50px;
    margin-bottom: 20px;
    padding: 10px 25px;
    border: 1px solid rgba(0, 0, 0, 0.1);
    border-radius: 20px;
    .new-item {
      width: 200px;
    }
    .home-footer {
      margin: 20px 0px;
    }
  }
}
.el-aside {
  min-height: 550px;
}
.messageSet {
  width: 550px;
  font-size: 12px;
  line-height: 40px;
  margin-bottom: 45px;
  ul {
    padding-top: 20px;
  }
  li {
    width: 100%;
    line-height: 40px;
  }
  li:nth-child(1) {
    .meLeft {
      border: none;
    }
  }
  li::after {
    content: "";
    width: 0;
    height: 0;
    display: block;
    clear: both;
  }
  li {
    zoom: 1;
  }
  .meLeft {
    float: left;
    width: 450px;
    color: #a9a9a9;
    cursor: pointer;
    border-top: 1px solid #a9a9a9;
    span {
      float: right;
    }
    p {
      float: left;
      width: 350px;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }
  }
  .meRight {
    float: right;
    color: #0000ff;
    cursor: pointer;
  }
  .moreMessage {
    cursor: pointer;
    color: #0000ff;
  }
}
.container {
  height: 100%;
  width: 100%;
  margin: 0 auto;
  .elcontainer {
    height: 100%;
  }
  span.title_span {
    margin-right: 21px;
  }
  h3.title_h3 {
    margin-top: 15px;
    font-size: 18px;
    font-family: PingFangSC-Regular;
    color: rgba(0, 0, 0, 1);
    font-weight: normal;
    .addMessage {
      margin-left: 40px;
      font-size: 12px;
      color: #0000ff;
      cursor: pointer;
    }
  }
  li {
    .el-input {
      height: 30px;
      // box-sizing: border-box;
      .el-input__inner {
        height: 30px;
      }
    }
  }
  .setNumber {
    height: 36px;
    width: 100%;
    display: flex;
    line-height: 36px;
    margin-top: 18px;
    .setNumberCon {
      display: flex;
      display: -webkit-flex;
      .groupid {
        margin-left: 20px;
        .el-input {
          margin: 0 10px;
          width: 150px;
        }
      }
    }
    .el-input-number {
      width: 79px;
      height: 36px;
      opacity: 0.4829;
      border-radius: 3px;
      margin-right: 20px;
      .el-input {
        input {
          height: 36px;
          color: #000;
        }
      }
    }
    .el-button {
      width: 79px;
      height: 40px;
      background: rgba(240, 89, 65, 1);
      border-radius: 3px;
      font-size: 14px;
      font-family: MicrosoftYaHei;
      color: rgba(255, 255, 255, 1);
    }
  }

  .el-header,
  .el-footer {
    background-color: #b3c0d1;
    color: #333;
    text-align: center;
  }
  .el-header {
    display: flex;
    padding: 0;
    position: relative;
    background: rgba(255, 255, 255, 1);
    .logo {
      margin-left: 20px;
      height: 100%;
      width: 120px;
      // background: red;
      img {
        max-height: 70px;
        max-width: 120px;
        margin-top: 5px;
      }
    }
    .title {
      margin-left: 20px;
      margin-top: 24px;
      width: 300px;
      height: 22px;
      font-size: 16px;
      font-family: PingFang-SC-Medium;
      color: rgba(33, 33, 33, 1);
      line-height: 22px;
    }
    .rightName {
      position: absolute;
      cursor: pointer;
      right: 0;
      height: 70px;
      line-height: 70px;
      padding-right: 16px;
    }
    .el-dropdown {
      position: absolute;
      cursor: pointer;
      right: 0;
      height: 70px;
      line-height: 70px;
      padding-right: 16px;
      .userHead {
        margin-left: 10px;
        display: inline-block;
        width: 45px;
        height: 45px;
        vertical-align: middle;
        cursor: pointer;
        img {
          width: 100%;
          height: 100%;
          border-radius: 50%;
        }
      }
    }
  }
  .el-aside {
    background-color: #1d2b36;
    color: white;
    text-align: center;
    line-height: 200px;
    .aside {
      width: 130px;
      height: 40px;
      line-height: 40px;
      text-align: center;
      background-color: #394555;
    }
    // *********************************************************左侧导航************************************************
    .el-menu {
      li {
        width: 130px;
      }
      .el-submenu {
        .el-submenu__title {
          padding-left: 0px !important;
          .el-submenu__icon-arrow {
            right: 14px;
            margin-top: -5px;
          }
        }
        .el-menu-item-group {
          .el-menu-item-group__title {
            padding: 0;
          }
          .el-menu-item {
            min-width: 0;
            font-size: 12px;
            height: 35px;
            line-height: 35px;
            padding: 0 !important;
            display: flex;
            align-items: center;
            justify-content: space-between;
            span:nth-child(1) {
              width: 24px;
              //   height: 24px;
              display: flex;
              justify-content: center;
              i {
                display: block;
                margin: 0;
                text-align: center;
                &::before {
                  margin: 0;
                  text-align: center;
                }
                &.el-icon-sort-up {
                  // margin-left: -12px;
                }
              }
            }
            span:nth-child(3) {
              width: 24px;
            }
          }
        }
      }
    }
  }
  .integralSet {
    .block {
      margin-bottom: 20px;
      .slogen {
        width: 500px;
      }
    }
    .demonstration {
      float: left;
      height: 40px;
      line-height: 40px;
      width: 150px;
    }
  }
  .basicSet {
    .block {
      margin-bottom: 20px;
      .slogen {
        width: 500px;
      }
    }
    .demonstration {
      float: left;
      height: 40px;
      line-height: 40px;
      width: 150px;
    }

    .setLogo {
      overflow: hidden;
      position: relative;
      margin-bottom: 30px;
      .uploadText {
        position: absolute;
        left: 345px;
        width: 300px;
        color: #f05941;
      }
      div.upload {
        float: left;
        width: 72px;
        height: 64px;
        margin-left: 60px;
        position: relative;
        input {
          position: absolute;
          bottom: 0;
          z-index: 99999;
          opacity: 0;
        }
        span.upload {
          position: absolute;
          bottom: 0;
          width: 100%;
          height: 25px;
          line-height: 30px;
          text-align: center;
          color: blue;
          cursor: pointer;
          font-size: 12px;
        }
      }
      .logoPic {
        float: left;
        height: 70px;
        width: 120px;
        border: 1px solid gray;
      }
      input {
        width: 72px;
      }
    }
    .setTpl {
      overflow: hidden;
      margin-bottom: 100px;
      img {
        position: relative;
      }
      .el-input {
        position: absolute;
        margin-top: 200px;
        margin-left: 50px;
        width: 100px;
      }
      div.uploadTpl {
        float: left;
        width: 72px;
        height: 64px;
        margin-left: 60px;
        margin-top: 240px;
        position: relative;
        input {
          position: absolute;
          bottom: 0;
          z-index: 99999;
          opacity: 0;
        }
        span.upload {
          position: absolute;
          bottom: 0;
          width: 100%;
          height: 25px;
          line-height: 30px;
          text-align: center;
          color: blue;
          cursor: pointer;
          font-size: 12px;
        }
      }
      .tplPic {
        float: left;
        height: 300px;
        width: 200px;
        border: 1px solid gray;
      }
      input {
        width: 100px;
      }
    }
    .classification {
      margin-bottom: 30px;
      .classificationTitle {
        font-size: 16px;
        margin-bottom: 20px;
      }
      .classificationText {
        font-size: 14px;
        margin-bottom: 20px;
      }
    }
  }

  .el-main {
    // background-color: #E9EEF3;
    text-align: center;
    box-sizing: border-box;
    padding: 23px 30px 0 30px;
    font-size: 14px;
    font-family: PingFangSC-Regular;
    color: rgba(0, 0, 0, 1);
    text-align: left;
    background-color: #fff;
    .channerName {
      height: 40px;
      width: 100%;
      // display:flex;
      .el-input {
        height: 40px;
        width: 278px;
        margin-right: 29px;
      }
      .el-button {
        width: 79px;
        height: 40px;
        background: rgba(240, 89, 65, 1);
        border-radius: 3px;
        font-size: 14px;
        font-family: MicrosoftYaHei;
        color: rgba(255, 255, 255, 1);
      }
    }
    .modelChoose {
      margin-top: 46px;
      padding-bottom: 54px;
      border-bottom: 1px rgba(161, 159, 159, 0.2) solid;
      display: flex;
      ul {
        li {
          float: left;
          width: 146px;
          height: 220px;
          margin-right: 50px;
          margin-bottom: 40px;
          position: relative;
          box-sizing: border-box;
          // border:3px solid yellowgreen;
          cursor: pointer;
          i {
            width: 100%;
            height: 30px;
            line-height: 30px;
            text-align: center;
            position: absolute;
            bottom: -30px;
            font-size: 12px;
            font-family: PingFang-SC-Medium;
            color: rgba(117, 117, 117, 1);
            .el-button,
            input {
              float: right;
              width: 86px;
              height: 26px;
              margin-top: 2px;
              line-height: 26px;
              padding: 0 5px;
            }
            input {
              position: absolute;
              top: 0;
              right: 0;
              padding: 0;
              opacity: 0;
              cursor: pointer;
            }
          }
          img.webImg {
            width: 100%;
            height: 100%;
            position: absolute;
            top: 0;
            left: 0;
          }
          img.chooseIcon {
            height: 29px;
            width: 29px;
            position: absolute;
            right: -16px;
            top: -16px;
          }
        }
      }
    }
    .setSlideshow {
      padding-top: 20px;
      border-bottom: 1px rgba(161, 159, 159, 0.2) solid;
      .el-input-number.is-controls-right .el-input__inner {
        padding-left: 10px;
      }
      ul {
        margin-top: 18px;
        overflow: hidden;
        li {
          float: left;
          width: 320px;
          height: 287px;
          margin-right: 20px;
          margin-bottom: 20px;
          box-sizing: border-box;
          border: 1px solid #dedede;
          padding: 10px;
          img {
            display: block;
            width: 100%;
            height: 125px;
          }
          p {
            display: flex;
            height: 30px;
            width: 100%;
            justify-content: space-between;
            margin-top: 14px;
            span {
              line-height: 30px;
              font-size: 12px;
              color: #777777;
            }
            .el-input {
              // width: 238px;
            }
            &.warn {
              height: 30px;
              span.text {
                font-size: 12px;
                color: rgba(248, 73, 73, 1);
                font-family: PingFangSC-Regular;
                line-height: 30px;
              }
              .el-button {
                padding: 0px 20px;
                height: 30px;
                span {
                  color: #fff;
                  font-size: 14px;
                }
              }
            }
          }
        }
      }
    }
    .setRecommend {
      padding: 20px 0px 0px 0px;
      border-bottom: 1px rgba(161, 159, 159, 0.2) solid;
      .el-input-number.is-controls-right .el-input__inner {
        padding-left: 10px;
      }
      .uploadText {
        position: absolute;
        left: 485px;
        width: 300px;
        color: #f05941;
      }
      div.upload {
        float: left;
        width: 72px;
        height: 64px;
        margin-left: 60px;
        position: relative;
        input {
          position: absolute;
          bottom: 0;
          z-index: 99999;
          opacity: 0;
        }
        span.upload {
          position: absolute;
          bottom: 0;
          width: 100%;
          height: 25px;
          line-height: 30px;
          text-align: center;
          color: blue;
          cursor: pointer;
          font-size: 12px;
        }
      }

      ul {
        overflow: hidden;
        margin-top: 20px;
        li {
          float: left;
          padding: 10px;
          margin-right: 20px;
          margin-bottom: 20px;
          width: 236px;
          height: 242px;
          box-sizing: border-box;
          border: 1px solid #d7d7d7;
          img {
            height: 97px;
            width: 100%;
            display: block;
          }
          p {
            margin-top: 20px;
            &.imgTitle {
              display: flex;
              justify-content: space-between;
              height: 30px;
              width: 100%;
              span {
                line-height: 30px;
                font-size: 12px;
                color: #777777;
              }
              .el-input {
                // width: 180px;
              }
            }
            &.warn {
              width: 100%;
              font-size: 12px;
              color: rgba(248, 73, 73, 1);
            }
          }
          .el-button {
            padding: 0px 20px;
            height: 30px;
            margin-top: 7px;
            span {
              color: #fff;
              font-size: 14px;
            }
          }
        }
      }
    }
    .setKeywords {
      .el-tag + .el-tag {
        margin-left: 10px;
      }
      .button-new-tag {
        margin-left: 10px;
        height: 32px;
        line-height: 30px;
        padding-top: 0;
        padding-bottom: 0;
      }
      .input-new-tag {
        width: 140px;
        margin-left: 10px;
        vertical-align: bottom;
      }
    }
    .setTag {
      padding: 20px 0;
      border-bottom: 1px rgba(161, 159, 159, 0.2) solid;
      .setTitleOrder {
        .el-table {
          .el-table__body {
            .el-table__row {
              // background: #000;
              .cell {
                height: 36px;
                line-height: 36px;
              }
              td {
                .cell {
                  position: relative;
                  span {
                    height: 20px;
                    width: 20px;
                    position: absolute;
                    img {
                      height: 20px;
                      width: 12px;
                      position: absolute;
                      top: 0;
                      left: 50%;
                      margin-left: -6px;
                    }
                    &:nth-child(1) {
                      top: 9px;
                      left: 58px;
                      cursor: pointer;
                    }
                    &:nth-child(2) {
                      top: 9px;
                      right: 58px;
                      cursor: pointer;
                    }
                    &:nth-child(3) {
                      display: none;
                      height: 0;
                      border: 1px dashed #777777;
                      cursor: disabled;
                    }
                  }
                }
              }
              &:nth-child(1) {
                .cell {
                  span:nth-child(2) {
                    visibility: hidden;
                  }
                  span:last-child {
                    display: block;
                    top: 18px;
                    right: 58px;
                  }
                }
              }
              &:last-child {
                .cell {
                  span:nth-child(1) {
                    visibility: hidden;
                  }
                  span:last-child {
                    display: block;
                    top: 18px;
                    left: 58px;
                  }
                }
              }
            }
          }
          margin-top: 18px;
          tr {
            height: 36px;
          }
          td,
          th {
            box-sizing: border-box;
            border: 1px solid #cdcdcd;
            height: 36px;
            padding: 0;
            text-align: center;
            color: #000;
            font-size: 12px;
            font-family: PingFangSC-Regular;
          }
          th {
            background-color: #d7d7d7;
          }
        }
      }
    }
    .BtnStyle {
      // border-radius: 20px!important;
      // margin: 0 auto!important;
    }
    .setShowType {
      margin-bottom: 20px;
      .el-input {
        width: 180px;
      }
    }
    .setBtn {
      padding: 40px 0px 50px 0;
      .el-button {
        margin-right: 90px;
        width: 130px;
        height: 40px;
        border-radius: 3px;
        &:nth-child(3) {
          background-color: #f05941;
          color: #fff;
          border-color: #f05941;
        }
      }
    }
  }
  .el-footer {
    background: rgba(24, 26, 26, 1);
  }

  body > .el-container {
    margin-bottom: 40px;
    height: 100%;
  }

  .el-container .el-aside {
    overflow: hidden;
    // height: 100%;
  }
  .setDialog {
    p {
      line-height: 40px;
      margin: 10px 0;
      &::after {
        content: "";
        width: 0;
        height: 0;
        display: block;
        clear: both;
      }
      .el-input {
        float: right;
        width: 450px;
      }
      .el-textarea {
        float: right;
        width: 450px;
        height: 150px;
        .el-textarea__inner {
          height: 100%;
        }
      }
    }
  }
}
</style>
