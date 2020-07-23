var gulp         = require('gulp');
var sass         = require('gulp-sass');
var uglify       = require('gulp-uglify');
var rename       = require('gulp-rename');
var include      = require('gulp-include');
var cleanCSS     = require('gulp-clean-css');
var autoprefixer = require('gulp-autoprefixer');
var merge        = require('merge-stream');

gulp.task('css', function() {
    return gulp.src('source/scss/main.scss')
        .pipe(sass({outputStyle:'expanded'}).on('error', sass.logError))
        .pipe(autoprefixer({
            browsers: ['last 2 versions'],
            cascade: false
        }))
        .pipe(cleanCSS())
        .pipe(rename({suffix:'.min'}))
        .pipe(gulp.dest('../asset/css'));
});

gulp.task('js', function() {
    return gulp.src('source/js/script.js')
        .pipe(include())
        .pipe(uglify())
        .pipe(rename({suffix:'.min'}))
        .pipe(gulp.dest('../asset/js'));
});

gulp.task('plugins', function() {
    return merge(
    gulp.src(['node_modules/@fortawesome/fontawesome-free/webfonts/*']).pipe(gulp.dest('../asset/font/fa')),
    gulp.src(['node_modules/dropzone/dist/*']).pipe(gulp.dest('../asset/plugin/dropzone')),
    gulp.src(['node_modules/dropzone/dist/min/*']).pipe(gulp.dest('../asset/plugin/dropzone/min')),
    gulp.src(['node_modules/select2/dist/css/*']).pipe(gulp.dest('../asset/plugin/select2/css')),
    gulp.src(['node_modules/select2/dist/js/*']).pipe(gulp.dest('../asset/plugin/select2/js')),
    gulp.src(['node_modules/datatables.net/js/*']).pipe(gulp.dest('../asset/plugin/datatables.net')),
    gulp.src(['node_modules/datatables.net-dt/*']).pipe(gulp.dest('../asset/plugin/datatables.net-dt')),
    gulp.src(['node_modules/datatables.net-dt/css/*']).pipe(gulp.dest('../asset/plugin/datatables.net-dt/css')),
    gulp.src(['node_modules/datatables.net-dt/images/*']).pipe(gulp.dest('../asset/plugin/datatables.net-dt/images')),
    gulp.src(['node_modules/datatables.net-dt/js/*']).pipe(gulp.dest('../asset/plugin/datatables.net-dt/js')),
    gulp.src(['node_modules/moment/locale/*']).pipe(gulp.dest('../asset/plugin/moment/locale')),
    gulp.src(['node_modules/moment/min/*']).pipe(gulp.dest('../asset/plugin/moment')),
    gulp.src(['node_modules/select2-bootstrap-theme/dist/*']).pipe(gulp.dest('../asset/plugin/select2-bootstrap-theme')),
    gulp.src(['node_modules/bootstrap-daterangepicker/*']).pipe(gulp.dest('../asset/plugin/bootstrap-daterangepicker')),
    gulp.src(['node_modules/perfect-scrollbar/css/*']).pipe(gulp.dest('../asset/plugin/perfect-scrollbar/css')),
    )


});

gulp.task('watch', function() {
    gulp.watch('source/scss/*.scss', ['css']);
    gulp.watch('source/js/*.js', ['js']);
});