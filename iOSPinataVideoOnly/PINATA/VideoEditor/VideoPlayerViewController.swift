import MobileCoreServices
import SCRecorder
import AVFoundation
import Photos


class VideoPlayerViewController: UIViewController {
    
    
    var canvasView: UIImageView = {
        let canvas = UIImageView()
        canvas.translatesAutoresizingMaskIntoConstraints = false
        return canvas
    }()
    
    var playerView: SCVideoPlayerView?
    var playerTimeObserver: Any?
    var captureImageView: UIImageView?
    
    var stickersViewController: StickersViewController = {
        let stickers = StickersViewController()
        stickers.view.translatesAutoresizingMaskIntoConstraints = false
        return stickers
    }()
    
    internal var exportSession: SCAssetExportSession?
    internal var asset: AVAsset?
    internal var outputUrl: URL?
    internal var completionRatios: [CGFloat] = []
    internal var cleanedCompletionRatios: [Double] = []
    internal var style: String?
    internal var capturedImage: UIImage?
    
    
    
    var colorRotation: [UIColor] = [.green, .blue, .red,  .magenta, .yellow, .cyan, .brown, .purple, .orange]
    
    var completionRatiosLayedOut: Bool = false
    var playerViewLayerLayedOut: Bool = false
    var assetIsMuted: Bool = false
    
    let stickerFullView: CGFloat = 100 // remainder of screen height
    var stickerPartialView: CGFloat {
        return UIScreen.main.bounds.height - 380
    }
    
    var stickersHeightContraint: NSLayoutConstraint?
    var sliderInvisibleThumbXConstraint: NSLayoutConstraint?
    
    var colorCollectionView: UICollectionView = {
        let layout: UICollectionViewFlowLayout = UICollectionViewFlowLayout()
        
    
        
        layout.scrollDirection = .horizontal
        layout.minimumInteritemSpacing = 0
        layout.minimumLineSpacing = UIScreen.main.bounds.width / (11 * 6)
        layout.itemSize = CGSize(width: (UIScreen.main.bounds.width / 11) - layout.minimumLineSpacing, height: (UIScreen.main.bounds.width / 11) - layout.minimumLineSpacing)
        let colorCollection = UICollectionView(frame: .zero, collectionViewLayout: layout)
        colorCollection.translatesAutoresizingMaskIntoConstraints = false
        colorCollection.backgroundColor = UIColor.clear
        colorCollection.isPagingEnabled = true
        return colorCollection
    }()
    
    var colorCollectionViewDelegate: ColorCollectionViewDelegate = {
        let colorCollectionViewDel = ColorCollectionViewDelegate()
        return colorCollectionViewDel
    }()
    
    
    var colorPickerView: UIView = {
        let colorPicker = UIView()
        colorPicker.translatesAutoresizingMaskIntoConstraints = false
        colorPicker.backgroundColor = .clear
        colorPicker.isHidden = true
        return colorPicker
    }()
    
    var colorPickerViewBottomConstraint: NSLayoutConstraint!
    
    var drawings : [UIImage?] = []
    var stickers : [UIImage] = []
    
    var stickersVCIsVisible = false
    var drawColor: UIColor = UIColor.white
    var textColor: UIColor = UIColor.white
    var backgroundTextColor: UIColor = UIColor.black.withAlphaComponent(0.5)
    var isDrawing: Bool = false
    var lastPoint: CGPoint = CGPoint(x: 0, y: 0)
    var swiped = false
    var lastPanPoint: CGPoint?
    var lastTextViewTransform: CGAffineTransform?
    var lastTextViewCenter: CGPoint?
    var lastTextViewParentBounds: CGRect?
    var lastTextViewFont:UIFont?
    
    
    var viewGestured: UIView?
    var textViewBoundsObserver: NSKeyValueObservation?
    
    var chaseTime = CMTime.zero
    var isSeekInProgress = false
    
    var isTyping: Bool = false
    
    var activeTextView: UITextView?
    
    var progressView: UIView = {
        let progress = UIView()
        progress.backgroundColor = UIColor(red: 60.0/255.0, green: 60.0/255.0, blue: 60.0/255.0, alpha: 0.6)
        progress.contentMode = .left
        progress.layer.cornerRadius = 15.0
        progress.translatesAutoresizingMaskIntoConstraints = false
        progress.clipsToBounds = true
        return progress
    }()
    
    var deleteView: UIImageView = {
        let delete = UIImageView(image: UIImage(named: "trash_icon")?.withAlignmentRectInsets(UIEdgeInsets(top: -5, left: -5, bottom: -5, right: -5)))

        delete.translatesAutoresizingMaskIntoConstraints = false
        delete.clipsToBounds = false
        delete.contentMode = .center
        delete.isHidden = true
        return delete
    }()
    
    var textButton: UIButton = {
        let text = UIButton(type: .custom)
        text.setImage(UIImage(named: "text_button"), for: .normal)
        text.translatesAutoresizingMaskIntoConstraints = false
        text.imageEdgeInsets = UIEdgeInsets(top: 5, left: 5, bottom: 5, right: 5)
        text.addTarget(self, action: #selector(handleTextButton(_:)), for: .touchUpInside)
        text.isHidden = false
        return text
    }()
    
    
    
 
    
    var drawButton: UIButton = {
        let draw = UIButton(type: .custom)
        //draw.setImage(UIImage(named: "text_button"), for: .normal)
        
        draw.translatesAutoresizingMaskIntoConstraints = false
        draw.setTitle("draw", for: .normal)

        draw.addTarget(self, action: #selector(handleDrawButton(_:)), for: .touchUpInside)
        draw.isHidden = false
        return draw
    }()
    
    var undoButton: UIButton = {
        let undo = UIButton(type: .custom)
        //clear.setImage(UIImage(named: "text_button"), for: .normal)
        
        undo.translatesAutoresizingMaskIntoConstraints = false
        undo.setTitle("undo", for: .normal)

        undo.addTarget(self, action: #selector(handleUndoButton(_:)), for: .touchUpInside)
        undo.isHidden = true
        return undo
    }()
    
    var stickerButton: UIButton = {
        let sticker = UIButton(type: .custom)
        //sticker.setImage(UIImage(named: "text_button"), for: .normal)
        
        sticker.translatesAutoresizingMaskIntoConstraints = false
        sticker.setTitle("sticker", for: .normal)

        sticker.addTarget(self, action: #selector(handleStickersButton(_:)), for: .touchUpInside)
        sticker.isHidden = false
        return sticker
    }()
    
    
    var cutButton: UIButton = {
        let cut = UIButton(type: .custom)
        //sticker.setImage(UIImage(named: "text_button"), for: .normal)
        
        cut.translatesAutoresizingMaskIntoConstraints = false
        cut.setTitle("cut", for: .normal)
        
        cut.addTarget(self, action: #selector(handleCutButton(_:)), for: .touchUpInside)
        cut.isHidden = true
        return cut
    }()
    
    var recordOverButton: UIButton = {
        let recordOver = UIButton(type: .custom)
        //sticker.setImage(UIImage(named: "text_button"), for: .normal)
        
        recordOver.translatesAutoresizingMaskIntoConstraints = false
        recordOver.setTitle("record", for: .normal)
        
        recordOver.addTarget(self, action: #selector(handleRecordOverButton(_:)), for: .touchUpInside)
        recordOver.isHidden = true
        return recordOver
    }()
    
    var saveButton: UIButton = {
        let save = UIButton(type: .custom)
        save.setImage(UIImage(named: "download_button"), for: .normal)
        save.setImage(UIImage(named: "download_complete_button"), for: .selected)
        save.translatesAutoresizingMaskIntoConstraints = false
        save.imageEdgeInsets = UIEdgeInsets(top: 5, left: 5, bottom: 5, right: 5)

        save.addTarget(self, action: #selector(handleSaveButton(_:)), for: .touchUpInside)
        save.isHidden = false
        return save
    }()
    
    var muteButton: UIButton = {
        let mute = UIButton(type: .custom)
        //sticker.setImage(UIImage(named: "text_button"), for: .normal)
        mute.setImage(UIImage(named: "speaker_button"), for: .normal)
        mute.setImage(UIImage(named: "mute_button"), for: .selected)
        mute.translatesAutoresizingMaskIntoConstraints = false
        mute.imageEdgeInsets = UIEdgeInsets(top: 5, left: 5, bottom: 5, right: 5)
        mute.addTarget(self, action: #selector(handleMuteButton(_:)), for: .touchUpInside)
        mute.isHidden = false
        return mute
    }()
    
    var doneButton: UIButton = {
        let done = UIButton(type: .custom)
        done.setImage(UIImage(named: "back_button"), for: .normal)
        done.translatesAutoresizingMaskIntoConstraints = false
        done.addTarget(self, action: #selector(handleDoneButton(_:)), for: .touchUpInside)
        done.isHidden = true
        return done
    }()
    
    var videoSlider: UISlider = {
        let slider = UISlider()
        slider.translatesAutoresizingMaskIntoConstraints = false
        slider.minimumTrackTintColor = UIColor.clear
        slider.maximumTrackTintColor = UIColor.clear
        slider.thumbTintColor = UIColor.white
        slider.backgroundColor = UIColor.clear
        
        
        slider.addTarget(self, action: #selector(handleVideoSliderChange), for: .valueChanged)
        
        return slider
    }()
    
    var sliderInvisibleThumb: UIView = {
        let thumb = UIView()
        thumb.translatesAutoresizingMaskIntoConstraints = false
        thumb.backgroundColor = UIColor.clear
        return thumb
    }()
    
    
    
    var cancelButton: UIButton = {
        let cancel = UIButton(type: .custom)
        cancel.translatesAutoresizingMaskIntoConstraints = false
        cancel.setImage(UIImage(named: "cancel_button"), for: .normal)
        cancel.addTarget(self, action: #selector(handleCancelButton(_:)), for: .touchUpInside)
        cancel.isHidden = false
        return cancel
    }()
    
    var completeButton: UIButton = {
        let complete = UIButton(type: .custom)
        complete.translatesAutoresizingMaskIntoConstraints = false
        complete.setImage(UIImage(named: "checkmark_icon"), for: .normal)
        complete.imageEdgeInsets = UIEdgeInsets(top: 10, left: 10, bottom: 10, right: 10)
        complete.addTarget(self, action: #selector(handleCompleteButton(_:)), for: .touchUpInside)
        complete.backgroundColor = UIColor(red: 210.0/255.0, green: 40/255.0, blue: 40.0/255.0, alpha: 1.0)
        
        complete.layer.cornerRadius = 30.0
        
        complete.isHidden = false
        return complete
    }()
    
    convenience init(asset: AVAsset, outputUrl: URL, completionRatios: [CGFloat], colorRotation: [UIColor], style: String) {
        self.init()
        self.asset = asset
        self.outputUrl = outputUrl
        self.completionRatios = completionRatios
        self.colorRotation = colorRotation
        self.style = style
    }
    
    convenience init(image: UIImage, style: String) {
        self.init()
        self.capturedImage = image
        self.style = style
    }
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        self.navigationItem.backBarButtonItem = UIBarButtonItem(title: "", style: .plain, target: nil, action: nil)
        self.navigationItem.backBarButtonItem?.tintColor = UIColor.white
        
        colorCollectionViewDelegate.colorDelegate = self
        
        colorCollectionView.delegate = colorCollectionViewDelegate
        colorCollectionView.dataSource = colorCollectionViewDelegate
        
        colorCollectionView.register(ColorCollectionViewCell.self, forCellWithReuseIdentifier: "ColorCollectionViewCell")
        
        
        stickersViewController.setStickersViewControllerDelegate(stickersDelegate: self)
        
        for i in 0...1 {
            self.stickers.append(UIImage(named: i.description )!)
        }
        
        for image in self.stickers {
            stickersViewController.stickers.append(image)
        }
        stickersViewController.collectionView.reloadData()
        
        self.addChild(stickersViewController)
        self.view.addSubview(stickersViewController.view)
        
        self.stickersHeightContraint = stickersViewController.view.heightAnchor.constraint(equalToConstant: 0)
        
        
     

        if let asset = self.asset {
            
            let playerItem = AVPlayerItem(asset: asset)
            let player = SCPlayer(playerItem: playerItem)
            player.autoRotate = false
            player.loopEnabled = true
            
            
            
            
            let playerView = SCVideoPlayerView(player: player)
            playerView.backgroundColor = .clear
            playerView.player = player
            playerView.delegate = self
            playerView.translatesAutoresizingMaskIntoConstraints = false

            self.playerView = playerView
            
        
            
            
            self.view.addSubview(self.playerView!)
            
            self.view.addSubview(progressView)
            
            
            for index in 0...completionRatios.count {
                let newView = UIView()
                newView.translatesAutoresizingMaskIntoConstraints = false
                self.progressView.addSubview(newView)
                newView.backgroundColor = colorRotation[index % colorRotation.count]
                NSLayoutConstraint.activate([
                    newView.heightAnchor.constraint(equalTo: self.progressView.heightAnchor),
                    newView.topAnchor.constraint(equalTo: self.progressView.topAnchor),
                    newView.bottomAnchor.constraint(equalTo: self.progressView.bottomAnchor),
                    ])
            }
            
            self.progressView.addSubview(self.videoSlider)
            
            
            
            self.view.addSubview(canvasView)
            addGestures(view: canvasView)
            
            
            
            
            self.view.addSubview(self.doneButton)
            self.view.addSubview(self.textButton)
            self.view.addSubview(self.completeButton)
            self.view.addSubview(self.saveButton)
            self.view.addSubview(self.deleteView)
            self.view.addSubview(self.colorPickerView)
            self.colorPickerView.addSubview(self.colorCollectionView)
            
            
            
            
            self.view.addSubview(self.cancelButton)
            self.view.addSubview(self.muteButton)
            self.view.addSubview(self.sliderInvisibleThumb)
            
            let longPress = UILongPressGestureRecognizer(target: self,
                                                    action: #selector(sliderInvisibleThumbLongPress))
            longPress.minimumPressDuration = 0.1
            longPress.delegate = self
            sliderInvisibleThumb.addGestureRecognizer(longPress)
            
            
            self.sliderInvisibleThumbXConstraint = self.sliderInvisibleThumb.leadingAnchor.constraint(equalTo: self.videoSlider.leadingAnchor)
            NSLayoutConstraint.activate([
                self.playerView!.topAnchor.constraint(equalTo: self.view.topAnchor),
                self.playerView!.bottomAnchor.constraint(equalTo: self.view.bottomAnchor),
                self.playerView!.leadingAnchor.constraint(equalTo: self.view.leadingAnchor),
                self.playerView!.trailingAnchor.constraint(equalTo: self.view.trailingAnchor),
                self.playerView!.widthAnchor.constraint(equalTo: self.view.widthAnchor),
                self.playerView!.heightAnchor.constraint(equalTo: self.view.heightAnchor),
                
                self.progressView.topAnchor.constraint(equalTo: self.view.topAnchor, constant: 30),
                self.progressView.widthAnchor.constraint(equalTo: self.view.widthAnchor, constant: -30),
                self.progressView.heightAnchor.constraint(equalToConstant: 30),
                self.progressView.centerXAnchor.constraint(equalTo: self.view.centerXAnchor),
                
                self.videoSlider.leadingAnchor.constraint(equalTo: self.progressView.leadingAnchor),
                self.videoSlider.trailingAnchor.constraint(equalTo: self.progressView.trailingAnchor),
                self.videoSlider.heightAnchor.constraint(equalTo: self.progressView.heightAnchor),
//                self.videoSlider.centerXAnchor.constraint(equalTo: self.progressView.centerXAnchor),
                self.videoSlider.centerYAnchor.constraint(equalTo: self.progressView.centerYAnchor),
                
                //width constant of 2 to match elliptical thumb
                self.sliderInvisibleThumb.widthAnchor.constraint(equalTo: self.progressView.heightAnchor),
                self.sliderInvisibleThumb.heightAnchor.constraint(equalTo: self.progressView.heightAnchor),
                self.sliderInvisibleThumb.centerYAnchor.constraint(equalTo: self.progressView.centerYAnchor),
                self.sliderInvisibleThumbXConstraint!,
                
                
                
                
                
                self.cancelButton.widthAnchor.constraint(equalToConstant: 40),
                self.cancelButton.heightAnchor.constraint(equalToConstant: 40),
                self.cancelButton.leadingAnchor.constraint(equalTo: self.view.leadingAnchor, constant: 25),
                self.cancelButton.topAnchor.constraint(equalTo: self.progressView.bottomAnchor, constant: 40),
                
                
                self.canvasView.topAnchor.constraint(equalTo: self.view.topAnchor),
                self.canvasView.bottomAnchor.constraint(equalTo: self.view.bottomAnchor),
                self.canvasView.leadingAnchor.constraint(equalTo: self.view.leadingAnchor),
                self.canvasView.trailingAnchor.constraint(equalTo: self.view.trailingAnchor),
                self.canvasView.heightAnchor.constraint(equalTo: self.view.heightAnchor),
                self.canvasView.widthAnchor.constraint(equalTo: self.view.widthAnchor),
                
                self.muteButton.widthAnchor.constraint(equalTo: self.cancelButton.widthAnchor),
                self.muteButton.heightAnchor.constraint(equalTo: self.cancelButton.heightAnchor),
                self.muteButton.leadingAnchor.constraint(equalTo: self.textButton.trailingAnchor, constant: 30),
                self.muteButton.bottomAnchor.constraint(equalTo: self.completeButton.bottomAnchor),
                
                self.saveButton.widthAnchor.constraint(equalTo: self.cancelButton.widthAnchor),
                self.saveButton.heightAnchor.constraint(equalTo: self.cancelButton.heightAnchor),
                self.saveButton.leadingAnchor.constraint(equalTo: self.muteButton.trailingAnchor, constant: 30),
                self.saveButton.bottomAnchor.constraint(equalTo: self.completeButton.bottomAnchor),

                ])
            
        } else if let image = self.capturedImage {
            
            self.captureImageView = UIImageView(image: image)
            self.captureImageView!.translatesAutoresizingMaskIntoConstraints = false
            self.captureImageView!.contentMode = .scaleAspectFit
            self.view.addSubview(self.captureImageView!)
            
            self.view.addSubview(canvasView)
            addGestures(view: canvasView)
            
            
            self.view.addSubview(self.doneButton)
            self.view.addSubview(self.textButton)
            self.view.addSubview(self.completeButton)
            self.view.addSubview(self.saveButton)
            self.view.addSubview(self.deleteView)
            self.view.addSubview(self.colorPickerView)
            self.colorPickerView.addSubview(self.colorCollectionView)
            
            self.view.addSubview(self.cancelButton)
            
            
            NSLayoutConstraint.activate([
                self.captureImageView!.topAnchor.constraint(equalTo: self.view.topAnchor),
                self.captureImageView!.bottomAnchor.constraint(equalTo: self.view.bottomAnchor),
                self.captureImageView!.leadingAnchor.constraint(equalTo: self.view.leadingAnchor),
                self.captureImageView!.trailingAnchor.constraint(equalTo: self.view.trailingAnchor),
                self.captureImageView!.widthAnchor.constraint(equalTo: self.view.widthAnchor),
                self.captureImageView!.heightAnchor.constraint(equalTo: self.view.heightAnchor),
                
                self.cancelButton.widthAnchor.constraint(equalToConstant: 40),
                self.cancelButton.heightAnchor.constraint(equalToConstant: 40),
                self.cancelButton.leadingAnchor.constraint(equalTo: self.view.leadingAnchor, constant: 25),
                self.cancelButton.topAnchor.constraint(equalTo: self.view.topAnchor, constant: 30),
                
                self.saveButton.widthAnchor.constraint(equalTo: self.cancelButton.widthAnchor),
                self.saveButton.heightAnchor.constraint(equalTo: self.cancelButton.heightAnchor),
                self.saveButton.leadingAnchor.constraint(equalTo: self.textButton.trailingAnchor, constant: 30),
                self.saveButton.bottomAnchor.constraint(equalTo: self.textButton.bottomAnchor),
                
                self.canvasView.topAnchor.constraint(equalTo: self.view.topAnchor),
                self.canvasView.bottomAnchor.constraint(equalTo: self.view.bottomAnchor),
                self.canvasView.leadingAnchor.constraint(equalTo: self.view.leadingAnchor),
                self.canvasView.trailingAnchor.constraint(equalTo: self.view.trailingAnchor),
                self.canvasView.heightAnchor.constraint(equalTo: self.view.heightAnchor),
                self.canvasView.widthAnchor.constraint(equalTo: self.view.widthAnchor),
                
                ])
        }
        

        
        //activated below
        self.colorPickerViewBottomConstraint = self.colorPickerView.bottomAnchor.constraint(equalTo: self.view.bottomAnchor, constant: 0)
        
        NSLayoutConstraint.activate([
            
            self.stickersViewController.view.bottomAnchor.constraint(equalTo: self.view.bottomAnchor),
            self.stickersViewController.view.leadingAnchor.constraint(equalTo: self.view.leadingAnchor),
            self.stickersViewController.view.trailingAnchor.constraint(equalTo: self.view.trailingAnchor),
            self.stickersHeightContraint!,
            
            
            
            self.doneButton.widthAnchor.constraint(equalTo: self.cancelButton.widthAnchor),
            self.doneButton.heightAnchor.constraint(equalTo: self.cancelButton.heightAnchor),
            self.doneButton.leadingAnchor.constraint(equalTo: self.cancelButton.leadingAnchor),
            self.doneButton.trailingAnchor.constraint(equalTo: self.cancelButton.trailingAnchor),
            self.doneButton.bottomAnchor.constraint(equalTo: self.cancelButton.bottomAnchor),
            
            
            
            self.completeButton.widthAnchor.constraint(equalToConstant: 60),
            self.completeButton.heightAnchor.constraint(equalToConstant: 60),
            self.completeButton.trailingAnchor.constraint(equalTo: self.view.trailingAnchor, constant: -25),
            self.completeButton.bottomAnchor.constraint(equalTo: self.view.bottomAnchor, constant: -15),
            
            
            self.deleteView.widthAnchor.constraint(equalTo: self.cancelButton.widthAnchor),
            self.deleteView.heightAnchor.constraint(equalTo: self.cancelButton.heightAnchor),
            self.deleteView.leadingAnchor.constraint(equalTo: self.cancelButton.leadingAnchor),
            self.deleteView.bottomAnchor.constraint(equalTo: self.completeButton.bottomAnchor),
            
            
            self.textButton.widthAnchor.constraint(equalTo: self.cancelButton.widthAnchor),
            self.textButton.heightAnchor.constraint(equalTo: self.cancelButton.heightAnchor),
            self.textButton.leadingAnchor.constraint(equalTo: self.cancelButton.leadingAnchor),
            self.textButton.bottomAnchor.constraint(equalTo: self.completeButton.bottomAnchor),
            
            self.colorPickerView.leadingAnchor.constraint(equalTo: self.view.leadingAnchor),
            self.colorPickerView.trailingAnchor.constraint(equalTo: self.view.trailingAnchor),
            self.colorPickerViewBottomConstraint,
            self.colorPickerView.heightAnchor.constraint(equalToConstant: 50),
            
            self.colorCollectionView.centerYAnchor.constraint(equalTo: self.colorPickerView.centerYAnchor),
            self.colorCollectionView.centerXAnchor.constraint(equalTo: self.colorPickerView.centerXAnchor),
            self.colorCollectionView.widthAnchor.constraint(equalTo: self.colorPickerView.widthAnchor),
            self.colorCollectionView.heightAnchor.constraint(equalTo: self.colorPickerView.heightAnchor),
            ])

        
        let edgePan = UIScreenEdgePanGestureRecognizer(target: self, action: #selector(screenEdgeSwiped))
        edgePan.edges = .bottom
        edgePan.delegate = self
        self.view.addGestureRecognizer(edgePan)
        
        
        
        
        let gesture = UIPanGestureRecognizer.init(target: self, action: #selector(stickersPanGesture))
        gesture.delegate = self
        stickersViewController.headerView.addGestureRecognizer(gesture)
        
       
        
    }
    
    override func viewDidLayoutSubviews() {
        super.viewDidLayoutSubviews()
        if !playerViewLayerLayedOut {
            if self.playerView != nil {
                let playLayer = AVPlayerLayer(player: self.playerView!.player)
                playLayer.frame = self.playerView!.bounds
                self.playerView!.layer.addSublayer(playLayer)
                self.playerView?.player?.play()
            }
            self.playerViewLayerLayedOut = true
        }
        
        if !completionRatiosLayedOut {
            for (index, completion) in completionRatios.enumerated() {
                var normalizedCompletion = (completion / completionRatios.last!)
                
                if index != 0 {
                    let newCompletion = max(0.0, completion - completionRatios[index - 1])
                    
                    
                    if normalizedCompletion > 0.0 && newCompletion > 0.0 {
                        //add to cleaned completion count
                        cleanedCompletionRatios.append(Double(normalizedCompletion))
                    }
                    
                    
                    normalizedCompletion = (newCompletion / completionRatios.last!)
                    progressView.subviews[index].widthAnchor.constraint(equalToConstant: normalizedCompletion * self.progressView.bounds.width).isActive = true
                    progressView.subviews[index].leadingAnchor.constraint(equalTo: progressView.subviews[index - 1].trailingAnchor).isActive = true
                } else {
                    
                    if normalizedCompletion > 0 {
                        //add to cleaned completion count
                        cleanedCompletionRatios.append(Double(normalizedCompletion))
                    }
                    
                    progressView.subviews[0].widthAnchor.constraint(equalToConstant: normalizedCompletion * self.progressView.bounds.width).isActive = true
                    progressView.subviews[0].leadingAnchor.constraint(equalTo: self.progressView.leadingAnchor).isActive = true
                }
                
            }
            self.completionRatiosLayedOut = true
            self.progressView.setNeedsLayout()
            
        }
        
    }
    


    override var prefersStatusBarHidden: Bool {
        return true
    }
    
    
    
    
    func shouldAutorotate() -> Bool {
        return false
    }
    
    override var supportedInterfaceOrientations: UIInterfaceOrientationMask {
        return [.portrait, .portraitUpsideDown]
    }
    
    
    override var preferredInterfaceOrientationForPresentation: UIInterfaceOrientation {
        return .portrait
    }
    
    override func viewWillAppear(_ animated: Bool) {
        super.viewWillAppear(animated)
        
        
        
        self.navigationController?.setNavigationBarHidden(true, animated: animated)
        
        NotificationCenter.default.addObserver(self,selector: #selector(keyboardWillShow), name: UIResponder.keyboardWillShowNotification, object: nil)
        NotificationCenter.default.addObserver(self, selector: #selector(keyboardDidShow), name: UIResponder.keyboardDidShowNotification, object: nil)
        NotificationCenter.default.addObserver(self, selector: #selector(keyboardWillChangeFrame(_:)), name: UIResponder.keyboardWillChangeFrameNotification, object: nil)
        NotificationCenter.default.addObserver(self, selector: #selector(keyboardWillHide), name: UIResponder.keyboardWillHideNotification, object: nil)
        NotificationCenter.default.addObserver(self,selector: #selector(keyboardDidHide), name: UIResponder.keyboardDidHideNotification, object: nil)
        
        
        
        
        // background event
        NotificationCenter.default.addObserver(self, selector: #selector(setPlayerLayerToNil), name: UIApplication.didEnterBackgroundNotification, object: nil)
        
        // foreground event
        NotificationCenter.default.addObserver(self, selector: #selector(reinitializePlayerLayer), name: UIApplication.willEnterForegroundNotification, object: nil)
        
        // add these 2 notifications to prevent freeze on long Home button press and back
        NotificationCenter.default.addObserver(self, selector: #selector(setPlayerLayerToNil), name: UIApplication.willResignActiveNotification, object: nil)
        
        NotificationCenter.default.addObserver(self, selector: #selector(reinitializePlayerLayer), name: UIApplication.didBecomeActiveNotification, object: nil)
        
        if self.playerView?.player != nil {
            self.addVideoTimeObserver()
            self.playerView?.player?.play()
        }
        
    }
    
    
    override func viewWillDisappear(_ animated: Bool) {
        super.viewWillDisappear(animated)
        self.navigationController?.setNavigationBarHidden(false, animated: animated)
        
        self.textViewBoundsObserver = nil
        
        
        NotificationCenter.default.removeObserver(self, name: UIResponder.keyboardWillShowNotification, object: nil)
        
        NotificationCenter.default.removeObserver(self, name: UIResponder.keyboardDidShowNotification, object: nil)
        
        NotificationCenter.default.removeObserver(self, name: UIResponder.keyboardWillChangeFrameNotification, object: nil)
        
        NotificationCenter.default.removeObserver(self, name: UIResponder.keyboardWillHideNotification, object: nil)
        
        NotificationCenter.default.removeObserver(self, name: UIResponder.keyboardDidHideNotification, object: nil)
        
        NotificationCenter.default.removeObserver(self, name: UIApplication.didEnterBackgroundNotification, object: nil)
        
        NotificationCenter.default.removeObserver(self, name: UIApplication.willEnterForegroundNotification, object: nil)
        
        NotificationCenter.default.removeObserver(self, name: UIApplication.willResignActiveNotification, object: nil)
        
        NotificationCenter.default.removeObserver(self, name: UIApplication.didBecomeActiveNotification, object: nil)
        
        if self.playerTimeObserver != nil {
            self.playerView?.player?.removeTimeObserver(self.playerTimeObserver!)
            self.playerTimeObserver = nil
        }
        
        
        
        self.playerView?.player?.pause()
        
        
    }
    
    deinit {
        
        print("PlayerViewDeinitialized")
    }
    
}

extension VideoPlayerViewController {
    // background event
    @objc fileprivate func setPlayerLayerToNil(){
        // first pause the player before setting the playerLayer to nil. The pause works similar to a stop button
        self.playerView?.player?.pause()
        var playerLayer = self.playerView?.playerLayer
        playerLayer?.removeFromSuperlayer()
        playerLayer = nil
    }
    
    // foreground event
    @objc fileprivate func reinitializePlayerLayer(){
        self.playerViewLayerLayedOut = false
        self.viewDidLayoutSubviews()
    }
}

extension VideoPlayerViewController: SCPlayerDelegate, SCVideoPlayerViewDelegate {
    func videoPlayerViewTapped(toPlay videoPlayerView: SCVideoPlayerView) {
        self.playerView?.player?.play()
    }
    
    func videoPlayerViewTapped(toPause videoPlayerView: SCVideoPlayerView) {
        self.playerView?.player?.pause()
    }
    
    
    
}

extension VideoPlayerViewController: ColorDelegate {
    
    // Observe ColorSlider .valueChanged events.
    @objc func changeColor(textColor: UIColor, backgroundColor: UIColor) {
        if isDrawing {
            if backgroundColor == UIColor.clear {
                self.drawColor = textColor
            } else {
                self.drawColor = backgroundColor
            }
        } else if activeTextView != nil {
//            (activeTextView?.subviews[0] as! UITextView).textColor = color
            activeTextView?.layer.backgroundColor = backgroundColor.cgColor
            activeTextView?.textColor = textColor
            self.textColor = textColor
            self.backgroundTextColor = backgroundColor
        }
      
    }
   
}

//_TEXTVIEW
extension VideoPlayerViewController: UITextViewDelegate {

    public func textViewDidChange(_ textView: UITextView) {
        let rotation = atan2(textView.superview?.transform.b ?? 0, textView.superview?.transform.a ?? 0)
        if rotation == 0 {
            let oldFrame = textView.superview?.frame
            let sizeToFit = textView.superview?.sizeThatFits(CGSize(width: oldFrame?.width ?? 0.0, height:CGFloat.greatestFiniteMagnitude))
            textView.superview?.frame.size = CGSize(width: oldFrame?.width ?? 0.0, height: sizeToFit?.height ?? 0.0)
        }
    }
    
    public func textViewDidBeginEditing(_ textView: UITextView) {
        isTyping = true
        lastTextViewTransform = textView.superview?.transform
        lastTextViewCenter = textView.superview?.center
        lastTextViewParentBounds = textView.superview?.bounds
        lastTextViewFont = textView.font!
        activeTextView = textView
        initializeTextViewBoundsObserver(view: textView)
        canvasView.bringSubviewToFront(textView.superview ?? textView)
        textView.font = UIFont(name: "Helvetica", size: 30)
        
        
        UIView.animate(withDuration: 0.3,
                       animations: {
                        textView.superview?.transform = CGAffineTransform.identity
                        textView.superview?.center = CGPoint(x: self.colorPickerView.center.x, y: self.colorPickerView.frame.origin.y - 30)
                       
        }, completion: nil)

    }

    public func textViewDidEndEditing(_ textView: UITextView) {
        guard lastTextViewTransform != nil && lastTextViewCenter != nil && lastTextViewFont != nil
            else {
                return
        }
        if textView.text.isEmpty {
            textView.superview?.removeFromSuperview()
            textView.removeFromSuperview()
        }
        self.saveButton.isSelected = false
        activeTextView = nil
        textView.superview?.bounds = self.lastTextViewParentBounds!
        textView.superview?.bounds.size.width = textView.bounds.width
        lastTextViewParentBounds = nil
        textView.font = self.lastTextViewFont!
        UIView.animate(withDuration: 0.3,
                       animations: {
                        textView.superview?.transform = self.lastTextViewTransform!
                        textView.superview?.center = self.lastTextViewCenter!
        }, completion: { [weak self] (complete) in
            if let view = textView.superview, complete {
                //self?.scaleEffect(view: view)
            }
        })
        
    }
    
}

//_KEYBOARD

extension VideoPlayerViewController {
    @objc func keyboardWillShow(notification: NSNotification) {
            doneButton.isHidden = false
            colorPickerView.isHidden = false
            cancelButton.isHidden = true
            self.view.bringSubviewToFront(colorPickerView)
    }
    
    @objc func keyboardDidShow(notification: NSNotification) {

    }
    
    @objc func keyboardDidHide(notification: NSNotification) {
        isTyping = false
    }
    
    @objc func keyboardWillHide(notification: NSNotification) {
        doneButton.isHidden = true
        colorPickerView.isHidden = true
        //hideToolbar(hide: false)
    }
    
    @objc func keyboardWillChangeFrame(_ notification: NSNotification) {
        if let userInfo = notification.userInfo {
            let endFrame = (userInfo[UIResponder.keyboardFrameEndUserInfoKey] as? NSValue)?.cgRectValue
            let duration:TimeInterval = (userInfo[UIResponder.keyboardAnimationDurationUserInfoKey] as? NSNumber)?.doubleValue ?? 0
            let animationCurveRawNSN = userInfo[UIResponder.keyboardAnimationCurveUserInfoKey] as? NSNumber
            let animationCurveRaw = animationCurveRawNSN?.uintValue ?? UIView.AnimationOptions.curveEaseInOut.rawValue
            let animationCurve:UIView.AnimationOptions = UIView.AnimationOptions(rawValue: animationCurveRaw)
            if (endFrame?.origin.y)! >= UIScreen.main.bounds.size.height {
                self.colorPickerViewBottomConstraint?.constant = 0.0
            } else {
                self.colorPickerViewBottomConstraint?.constant = -1 * (endFrame?.size.height ?? 0.0)
            }
            UIView.animate(withDuration: duration,
                           delay: TimeInterval(0),
                           options: animationCurve,
                           animations: { self.view.layoutIfNeeded() },
                           completion: nil)
        }
    }
    
}

//_STICKERS
extension VideoPlayerViewController {

    
    //MARK: Pan Gesture
    
    @objc func stickersPanGesture(_ recognizer: UIPanGestureRecognizer) {
        
        self.stickersViewController.view.setNeedsUpdateConstraints()
        
        let translation = recognizer.translation(in: self.view)
        let velocity = recognizer.velocity(in: self.view)
        
        let y = self.stickersViewController.view.frame.minY
        if y + translation.y >= self.stickerFullView {
            let newMinY = y + translation.y
            self.stickersHeightContraint?.constant = UIScreen.main.bounds.height - newMinY
            self.view.layoutIfNeeded()
            recognizer.setTranslation(CGPoint.zero, in: self.view)
        }
        
        if recognizer.state == .ended {
            var duration =  velocity.y < 0 ? Double((y - self.stickerFullView) / -velocity.y) : Double((self.stickerPartialView - y) / velocity.y )
            duration = duration > 1.3 ? 1 : duration
            //velocity is direction of gesture
            UIView.animate(withDuration: duration, delay: 0.0, options: [.allowUserInteraction], animations: {
                if  velocity.y >= 0 {
                    if y + translation.y >= self.stickerPartialView  {
                        self.removeStickersView()
                    } else {
                        self.stickersHeightContraint?.constant = UIScreen.main.bounds.height - self.stickerPartialView
                        self.view.layoutIfNeeded()
                    }
                } else {
                    if y + translation.y >= self.stickerPartialView  {
                        self.stickersHeightContraint?.constant = UIScreen.main.bounds.height - self.stickerPartialView
                        self.view.layoutIfNeeded()
                    } else {
                        self.stickersHeightContraint?.constant = UIScreen.main.bounds.height - self.stickerFullView
                        self.view.layoutIfNeeded()
                    }
                }
                
            }, completion: nil)
        }
    }
    
 
    func showStickersViewController() {
        self.stickersVCIsVisible = true
        //hideToolbar(hide: true)
        self.canvasView.isUserInteractionEnabled = false
        stickersViewController.didMove(toParent: self)
        
        self.stickersViewController.view.setNeedsUpdateConstraints()
        
        UIView.animate(withDuration: 0.6) { [weak self] in
            self?.stickersHeightContraint?.constant = UIScreen.main.bounds.height - (self?.stickerPartialView ?? UIScreen.main.bounds.height)
            self?.stickersViewController.view.layoutIfNeeded()
        }
    }
    
    func removeStickersView() {
        stickersVCIsVisible = false
        self.canvasView.isUserInteractionEnabled = true
        
        self.stickersViewController.view.setNeedsUpdateConstraints()

        UIView.animate(withDuration: 0.3,
                       delay: 0,
                       options: UIView.AnimationOptions.curveEaseOut,
                       animations: { () -> Void in
                        
                        self.stickersHeightContraint?.constant = 0
                        self.stickersViewController.view.layoutIfNeeded()

                        
        }, completion: { (finished) -> Void in
            
            //self.hideToolbar(hide: false)
        })
    }
}

extension VideoPlayerViewController: StickersViewControllerDelegate {
    
    func didSelectView(view: UIView) {
        self.removeStickersView()
        
        view.center = canvasView.center
        self.canvasView.addSubview(view)
        //Gestures
        addGestures(view: view)
    }
    
    func didSelectImage(image: UIImage) {
        self.removeStickersView()
        
        let imageView = UIImageView(image: image)
        imageView.contentMode = .scaleAspectFit
        imageView.frame.size = CGSize(width: 150, height: 150)
        imageView.center = canvasView.center
        
        self.canvasView.addSubview(imageView)
        //Gestures
        addGestures(view: imageView)
    }

    
    func addGestures(view: UIView) {
        //Gestures
        view.isUserInteractionEnabled = true
        
        let panGesture = UIPanGestureRecognizer(target: self,
                                                action: #selector(VideoPlayerViewController.panGesture))
        panGesture.minimumNumberOfTouches = 1
        panGesture.maximumNumberOfTouches = 1
        panGesture.delegate = self
        view.addGestureRecognizer(panGesture)
        
        let pinchGesture = UIPinchGestureRecognizer(target: self,
                                                    action: #selector(VideoPlayerViewController.pinchGesture))
        pinchGesture.delegate = self
        view.addGestureRecognizer(pinchGesture)
        
        let rotationGestureRecognizer = UIRotationGestureRecognizer(target: self,
                                                                    action:#selector(VideoPlayerViewController.rotationGesture) )
        rotationGestureRecognizer.delegate = self
        view.addGestureRecognizer(rotationGestureRecognizer)
        
        let tapGesture = UITapGestureRecognizer(target: self, action: #selector(VideoPlayerViewController.tapGesture))
        view.addGestureRecognizer(tapGesture)
        
    }
}

//video scrubbing
extension VideoPlayerViewController {
    
    func executeVideoSliderChange() {
        if let duration = self.playerView?.player?.currentItem?.duration {
            
            
            let seekTime = CMTime(seconds: Float64(videoSlider.value) * duration.seconds, preferredTimescale: 1000)
            self.stopPlayingAndSeekSmoothlyToTime(newChaseTime: seekTime)
            
//            self.playerView?.player?.seek(to: seekTime, completionHandler: { (completedSeek) in
//
//            })
        }
    }
    
    func addVideoTimeObserver() {
        if let duration = self.playerView?.player?.currentItem?.duration, duration.value != 0 {
            
            self.playerTimeObserver = self.playerView?.player?.addPeriodicTimeObserver(forInterval: CMTime(value: 1, timescale: CMTimeScale(100 * duration.value)), queue: DispatchQueue.main) { [weak self] (progressTime) in
                self?.sliderInvisibleThumbXConstraint?.constant = CGFloat(self?.videoSlider.value ?? 0.0) * ((self?.videoSlider.frame.width ?? 0.0) - (self?.videoSlider.frame.height ?? 0.0))
                self?.videoSlider.value = Float(progressTime.seconds / duration.seconds)
            }
            
        }
        
    }
    
    func stopPlayingAndSeekSmoothlyToTime(newChaseTime:CMTime)
    {
        
        if CMTimeCompare(newChaseTime, chaseTime) != 0
        {
            chaseTime = newChaseTime;
            
            if !isSeekInProgress
            {
                trySeekToChaseTime()
            }
        }
    }
    
    func trySeekToChaseTime()
    {
        if self.playerView?.player?.status == .unknown
        {
            // wait until item becomes ready (KVO player.currentItem.status)
        }
        else if self.playerView?.player?.status == .readyToPlay
        {
            actuallySeekToTime()
        }
    }
    
    func actuallySeekToTime()
    {
        isSeekInProgress = true
        let seekTimeInProgress = chaseTime
        self.playerView?.player?.seek(to: seekTimeInProgress, toleranceBefore: CMTime.zero, toleranceAfter: CMTime.zero, completionHandler:
            { [weak self] (isFinished:Bool) -> Void in
                
                if CMTimeCompare(seekTimeInProgress, self?.chaseTime ?? CMTime.invalid) == 0
                {
                    self?.isSeekInProgress = false
                }
                else
                {
                    self?.trySeekToChaseTime()
                }
        })
    }
    

}


//_GESTURES
extension VideoPlayerViewController : UIGestureRecognizerDelegate  {
    
    func findClosestSubview(onView view: UIView, toLocation location: CGPoint, paddingMultiplier: CGFloat?) -> UIView? {
        let subviews = view.subviews.compactMap{ $0 }
        let distances = subviews.map { $0.center.distance(to: location) }
        
        // we zip both sequences into one, that way we don't have to worry about sorting two arrays
        let subviewDistances = zip(subviews, distances)
        // we don't need sorting just to get the minimum
        let closestSubview = subviewDistances.min { $0.1 < $1.1 }
        if let closestSubview = closestSubview?.0 {
            var testHitFrame = closestSubview.frame
            if paddingMultiplier != nil {
                let newWidth = closestSubview.bounds.width * paddingMultiplier!
                let newHeight = closestSubview.bounds.height * paddingMultiplier!
                testHitFrame = CGRect(x: closestSubview.frame.origin.x - newWidth / 2, y: closestSubview.frame.origin.y - newHeight / 2, width: newWidth, height: newHeight)
            }
            if testHitFrame.contains(location) {
                return closestSubview
            }
        }
        return nil
    }
    
    func initializeGestureActiveView(recognizer: UIGestureRecognizer) {
        if viewGestured == nil && recognizer.view != nil {
            let location = recognizer.location(in: canvasView)
            let view = recognizer.view!
            
            //update Save Button for changes
            self.saveButton.isSelected = false
            
            //not the canvas
            if view != canvasView {
                viewGestured = view
                
            } else if view == canvasView {
                //it's the canvas
                if let closest = findClosestSubview(onView: canvasView, toLocation: location, paddingMultiplier: 3) {
                    viewGestured = closest
                }
            }
            
            initializeTextViewBoundsObserver(view: viewGestured)
        }
    }
    
    func initializeTextViewBoundsObserver(view: UIView?) {
        if view == self.viewGestured {
            for subview in view?.subviews ?? [] {
                if subview is UITextView {
                    textViewBoundsObserver = subview.observe(\.bounds, changeHandler: { [weak self] (viewGesturedObserved, change) in
                        view?.bounds.size = viewGesturedObserved.bounds.size
                    })
                }
            }
        } else if view is UITextView {
            textViewBoundsObserver = view?.observe(\.bounds, changeHandler: { [weak self](viewGesturedObserved, change) in
                view?.superview?.bounds.size = viewGesturedObserved.bounds.size
            })
        }
    }
    
    @objc func sliderInvisibleThumbLongPress(_ recognizer: UILongPressGestureRecognizer) {
        
        if recognizer.state == .began {
            self.playerView?.player?.pause()
        } else if recognizer.state == .ended {
           self.playerView?.player?.play()
        }
        if recognizer.state == .changed {
            let pressedLocation = recognizer.location(in: self.videoSlider)
            
            let fixedYLocation = CGPoint(x: pressedLocation.x, y: self.videoSlider.frame.height / 2)
            
            //let hitTestView = videoSlider.hitTest(fixedYLocation, with: nil)
            
            
            if videoSlider.frame.contains(fixedYLocation) {
                
                self.videoSlider.value = Float(fixedYLocation.x / videoSlider.frame.maxX)
                print(self.videoSlider.value)
                //self.videoSlider.setValue(Float(fixedYLocation.x / videoSlider.frame.maxX), animated: false)
                self.sliderInvisibleThumbXConstraint?.constant = CGFloat(self.videoSlider.value) * (self.videoSlider.frame.width - self.videoSlider.frame.height)
                
                executeVideoSliderChange()
                
            }
        }
        
        
    }
    
//    /**
//     UIPanGestureRecognizer - Moving Objects
//     Selecting transparent parts of the imageview won't move the object
//     */
//    @objc func panGesture(_ recognizer: UIPanGestureRecognizer) {
//        initializeGestureActiveView(recognizer: recognizer)
//        if viewGestured != nil {
//            moveView(view: viewGestured!, recognizer: recognizer)
//        }
//    }
//
//    /**
//     Moving Objects
//     delete the view if it's inside the delete view
//     Snap the view back if it's out of the canvas
//     */
//
//    func moveView(view: UIView, recognizer: UIPanGestureRecognizer)  {
//
//        //hideToolbar(hide: true)
//        deleteView.isHidden = false
//
//        view.superview?.bringSubviewToFront(view)
//        let pointToSuperView = recognizer.location(in: self.view)
//
//        view.center = CGPoint(x: view.center.x + recognizer.translation(in: canvasView).x,
//                              y: view.center.y + recognizer.translation(in: canvasView).y)
//
//        recognizer.setTranslation(CGPoint.zero, in: canvasView)
//
//        if let previousPoint = lastPanPoint {
//            //View is going into deleteView
//            if deleteView.frame.contains(pointToSuperView) && !deleteView.frame.contains(previousPoint) {
//                if #available(iOS 10.0, *) {
//                    let generator = UIImpactFeedbackGenerator(style: .heavy)
//                    generator.impactOccurred()
//                }
//                UIView.animate(withDuration: 0.3, animations: {
//                    view.transform = view.transform.scaledBy(x: 0.25, y: 0.25)
//                    view.center = recognizer.location(in: self.canvasView)
//                })
//            }
//                //View is going out of deleteView
//            else if deleteView.frame.contains(previousPoint) && !deleteView.frame.contains(pointToSuperView) {
//                //Scale to original Size
//                UIView.animate(withDuration: 0.3, animations: {
//                    view.transform = view.transform.scaledBy(x: 4, y: 4)
//                    view.center = recognizer.location(in: self.canvasView)
//                })
//            }
//        }
//        lastPanPoint = pointToSuperView
//
//        if recognizer.state == .ended {
//            viewGestured = nil
//            lastPanPoint = nil
//            //hideToolbar(hide: false)
//            deleteView.isHidden = true
//            let point = recognizer.location(in: self.view)
//
//            if deleteView.frame.contains(point) { // Delete the view
//                view.removeFromSuperview()
//                if #available(iOS 10.0, *) {
//                    let generator = UINotificationFeedbackGenerator()
//                    generator.notificationOccurred(.success)
//                }
//            } else if !canvasView.bounds.contains(view.center) { //Snap the view back to canvasView
//                UIView.animate(withDuration: 0.3, animations: {
//                    view.center = self.canvasView.center
//                })
//
//            }
//        }
//    }
    
    /**
     UIPanGestureRecognizer - Moving Objects
     Selecting transparent parts of the imageview won't move the object
     */
    @objc func panGesture(_ recognizer: UIPanGestureRecognizer) {
        if recognizer.state == .began {
            initializeGestureActiveView(recognizer: recognizer)
        }
        if self.viewGestured != nil && recognizer.view != self.canvasView {
            //hideToolbar(hide: true)
            self.deleteView.isHidden = false
            self.completeButton.isHidden = true
            self.textButton.isHidden = true
            self.muteButton.isHidden = true
            self.saveButton.isHidden = true
            
            self.viewGestured!.superview?.bringSubviewToFront(self.viewGestured!)
            let pointToSuperView = recognizer.location(in: self.view)
            

            self.viewGestured!.center = CGPoint(x: self.viewGestured!.center.x + recognizer.translation(in: self.canvasView).x,
                                  y: self.viewGestured!.center.y + recognizer.translation(in: self.canvasView).y)
            
            recognizer.setTranslation(CGPoint.zero, in: self.canvasView)
            
            
            
            if let previousPoint = self.lastPanPoint {
                //View is going into deleteView
                if self.deleteView.frame.contains(pointToSuperView) && !self.deleteView.frame.contains(previousPoint) {
                    if #available(iOS 10.0, *) {
                        let generator = UIImpactFeedbackGenerator(style: .heavy)
                        generator.impactOccurred()
                    }
                    UIView.animate(withDuration: 0.3, animations: {
                        self.viewGestured!.transform = self.viewGestured!.transform.scaledBy(x: 0.25, y: 0.25)
                        self.viewGestured!.center = recognizer.location(in: self.canvasView)
                    })
                }
                    //View is going out of deleteView
                else if deleteView.frame.contains(previousPoint) && !deleteView.frame.contains(pointToSuperView) {
                    //Scale to original Size
                    UIView.animate(withDuration: 0.3, animations: {
                        self.viewGestured!.transform = self.viewGestured!.transform.scaledBy(x: 4, y: 4)
                        self.viewGestured!.center = recognizer.location(in: self.canvasView)
                    })
                }
            }
            lastPanPoint = pointToSuperView
            
            if recognizer.state == .ended {
                let point = recognizer.location(in: self.view)
                
                if self.deleteView.frame.contains(point) { // Delete the view
                    self.viewGestured!.removeFromSuperview()
                    if #available(iOS 10.0, *) {
                        let generator = UINotificationFeedbackGenerator()
                        generator.notificationOccurred(.success)
                    }
                } else if !self.canvasView.bounds.contains(self.viewGestured!.center) {
                    //Snap the view back to canvasView
                    UIView.animate(withDuration: 0.3, animations: {
                        self.viewGestured!.center = self.canvasView.center
                    })
                    
                }
                
                if recognizer.state == .ended {
                    //hideToolbar(hide: false)
                    self.deleteView.isHidden = true
                    self.completeButton.isHidden = false
                    self.textButton.isHidden = false
                    self.muteButton.isHidden = false
                    self.saveButton.isHidden = false
                    self.viewGestured = nil
                    self.textViewBoundsObserver = nil
                    self.lastPanPoint = nil
                    
                }
            }
            
            
        }
        
    }
    
    /**
     UIPinchGestureRecognizer - Pinching Objects
     If it's a UITextView will make the font bigger so it doen't look pixlated
     */
    @objc func pinchGesture(_ recognizer: UIPinchGestureRecognizer) {
        if recognizer.state == .began {
            initializeGestureActiveView(recognizer: recognizer)
        } else if recognizer.state == .ended {
            if self.lastPanPoint == nil {
                viewGestured = nil
                textViewBoundsObserver = nil
            }
        }
        if viewGestured != nil {
            
            if !(viewGestured!.subviews.isEmpty) && viewGestured!.subviews.last is UITextView {
                let textView = viewGestured!.subviews.last as! UITextView
                if textView.font!.pointSize * recognizer.scale < 90 {
                    let font = UIFont(name: textView.font!.fontName, size: textView.font!.pointSize * recognizer.scale)
                    textView.font = font
                }
                
                let sizeToFit = textView.sizeThatFits(CGSize(width: UIScreen.main.bounds.size.width, height:CGFloat.greatestFiniteMagnitude))
                textView.bounds.size = CGSize(width: textView.intrinsicContentSize.width, height: sizeToFit.height)

                
                textView.setNeedsDisplay()
            }
            else {
                viewGestured!.transform = viewGestured!.transform.scaledBy(x: recognizer.scale, y: recognizer.scale)
            }
            recognizer.scale = 1
        }
    }
    
    /**
     UIRotationGestureRecognizer - Rotating Objects
     */
    @objc func rotationGesture(_ recognizer: UIRotationGestureRecognizer) {
        if recognizer.state == .began {
            initializeGestureActiveView(recognizer: recognizer)
        } else if recognizer.state == .ended {
            if self.lastPanPoint == nil {
                viewGestured = nil
                textViewBoundsObserver = nil
            }
            
        }
    
        if viewGestured != nil {
            viewGestured!.transform = viewGestured!.transform.rotated(by: recognizer.rotation)
            recognizer.rotation = 0
        }
    }
    
    /**
     UITapGestureRecognizer - Taping on Objects
     Will make scale scale Effect
     Selecting transparent parts of the imageview won't move the object
     */
    @objc func tapGesture(_ recognizer: UITapGestureRecognizer) {
        if let view = recognizer.view {
//            for subview in view.subviews {
//                if subview is UITextView {
//                    textViewDidBeginEditing(subview as! UITextView)
//                    return
//                }
//            }
            if view is UIImageView && view != self.canvasView {
                //Tap only on visible parts on the image
                for imageView in subImageViews(view: canvasView) {
                    let location = recognizer.location(in: imageView)
                    let alpha = imageView.alphaAtPoint(location)
                    if alpha > 0 {
                        scaleEffect(view: imageView)
                        break
                    }
                }
                
            } else if view == self.canvasView && self.activeTextView == nil && self.isTyping == false {
                executeTextButtonAction()
            } else if view != self.canvasView {
                scaleEffect(view: view)
            }
        }
    }
    
    /*
     Support Multiple Gesture at the same time
     */
    public func gestureRecognizer(_ gestureRecognizer: UIGestureRecognizer, shouldRecognizeSimultaneouslyWith otherGestureRecognizer: UIGestureRecognizer) -> Bool {
        return true
    }
    
    public func gestureRecognizer(_ gestureRecognizer: UIGestureRecognizer, shouldRequireFailureOf otherGestureRecognizer: UIGestureRecognizer) -> Bool {
        return false
    }
    
    public func gestureRecognizer(_ gestureRecognizer: UIGestureRecognizer, shouldBeRequiredToFailBy otherGestureRecognizer: UIGestureRecognizer) -> Bool {
        return false
    }
    
    @objc func screenEdgeSwiped(_ recognizer: UIScreenEdgePanGestureRecognizer) {
        if recognizer.state == .recognized {
            if !stickersVCIsVisible {
                showStickersViewController()
            }
        }
    }
    
    /**
     Scale Effect
     */
    func scaleEffect(view: UIView) {
        view.superview?.bringSubviewToFront(view)
        
        if #available(iOS 10.0, *) {
            let generator = UIImpactFeedbackGenerator(style: .heavy)
            generator.impactOccurred()
        }
        let previouTransform =  view.transform
        UIView.animate(withDuration: 0.2,
                       animations: {
                        view.transform = view.transform.scaledBy(x: 1.2, y: 1.2)
        },
                       completion: { _ in
                        UIView.animate(withDuration: 0.2) {
                            view.transform  = previouTransform
                        }
        })
    }
    
    /**
     Moving Objects
     delete the view if it's inside the delete view
     Snap the view back if it's out of the canvas
     */
    
    
    
    func subImageViews(view: UIView) -> [UIImageView] {
        var imageviews: [UIImageView] = []
        for imageView in view.subviews {
            if imageView is UIImageView {
                imageviews.append(imageView as! UIImageView)
            }
        }
        return imageviews
    }
}

//_DRAWING
//
extension VideoPlayerViewController {
    
    override public func touchesBegan(_ touches: Set<UITouch>,
                                      with event: UIEvent?){
        if isDrawing {
            swiped = false
            if let touch = touches.first {
                lastPoint = touch.location(in: self.canvasView)
            }
        }
            //Hide stickersVC if clicked outside it
        else if stickersVCIsVisible == true {
            if let touch = touches.first {
                let location = touch.location(in: self.view)
                if !stickersViewController.view.frame.contains(location) {
                    removeStickersView()
                }
            }
        }
        
        else if activeTextView != nil {
            if let touch = touches.first {
                let location = touch.location(in: self.view)
                if location.y < colorPickerView.frame.minY {
                    executeDoneAction()
                }
            }
        }
        
        
    }
    
    override public func touchesMoved(_ touches: Set<UITouch>,
                                      with event: UIEvent?){
        if isDrawing {
            // 6
            swiped = true
            if let touch = touches.first {
                let currentPoint = touch.location(in: canvasView)
                drawLineFrom(lastPoint, toPoint: currentPoint)
                
                // 7
                lastPoint = currentPoint
            }
        }
    }
    
    override public func touchesEnded(_ touches: Set<UITouch>,
                                      with event: UIEvent?){
        if isDrawing {
            if !swiped {
                // draw a single point
                drawLineFrom(lastPoint, toPoint: lastPoint)
            }
        }
        self.drawings.append(canvasView.image)
    }
    
    func drawLineFrom(_ fromPoint: CGPoint, toPoint: CGPoint) {
        // 1
        let canvasSize = canvasView.frame.integral.size
        UIGraphicsBeginImageContextWithOptions(canvasSize, false, 0)
        if let context = UIGraphicsGetCurrentContext() {
            canvasView.image?.draw(in: CGRect(x: 0, y: 0, width: canvasSize.width, height: canvasSize.height))
            // 2
            context.move(to: CGPoint(x: fromPoint.x, y: fromPoint.y))
            context.addLine(to: CGPoint(x: toPoint.x, y: toPoint.y))
            // 3
            context.setLineCap( CGLineCap.round)
            context.setLineWidth(5.0)
            context.setStrokeColor(drawColor.cgColor)
            context.setBlendMode( CGBlendMode.normal)
            // 4
            context.strokePath()
            // 5
            let newImage = UIGraphicsGetImageFromCurrentImageContext()
            
            canvasView.image = newImage
        }
        UIGraphicsEndImageContext()
    }
    
}


//_EPXPORTING

extension VideoPlayerViewController {
    internal func executeExport(completion: ((_ error: Error?, _ cancelled: Bool) -> ())?) {
        
            if self.exportSession != nil {
                self.exportSession?.cancelExport()
                self.exportSession = nil
            }
        
            self.exportSession = SCAssetExportSession(asset: self.asset!)
            exportSession?.videoConfiguration.preset = SCPresetHighestQuality
            exportSession?.audioConfiguration.preset = SCPresetHighestQuality
            exportSession?.videoConfiguration.maxFrameRate = CMTimeScale(35)
            let finalCanvasImage = self.canvasView.toImage()
            let finalCanvasImageView: UIImageView & SCVideoOverlay = CanvasView()
            finalCanvasImageView.frame = self.canvasView.frame
            finalCanvasImageView.image = finalCanvasImage
            exportSession?.videoConfiguration.overlay = finalCanvasImageView
        
//            exportSession.videoConfiguration.watermarkImage = self.canvasView.toImage()
//        exportSession.videoConfiguration.watermarkAnchorLocation = SCWatermarkAnchorLocationTopLeft
//        exportSession.videoConfiguration.watermarkAnchorLocation = SCWatermarkAnchorLocationBottomLeft
        
        
            if assetIsMuted {
                let audioTracks = asset?.tracks(withMediaType: .audio)
                var allAudioParams: [AVMutableAudioMixInputParameters] = []
                for track: AVAssetTrack in audioTracks ?? [] {
                    let audioInputParams = AVMutableAudioMixInputParameters()
                    audioInputParams.setVolume(0.0, at: .zero)
                    audioInputParams.trackID = track.trackID
                    allAudioParams.append(audioInputParams)
                }
                
                let zeroAudioMix = AVMutableAudioMix()
                zeroAudioMix.inputParameters = allAudioParams
                
                exportSession?.audioConfiguration.audioMix = zeroAudioMix
            }
            
            exportSession?.outputUrl = self.outputUrl
            exportSession?.outputFileType = AVFileType.mp4.rawValue
            exportSession?.contextType = SCContextType.auto
            exportSession?.delegate = self
            exportSession?.contextType = .auto
        
            DispatchQueue.main.async{ [weak self] in
                self?.exportSession?.exportAsynchronously(completionHandler: { [weak self] in
                    
                if let error = self?.exportSession?.error {
                    let alert = UIAlertController(title: "Something wrong with exporting video", message: "Please try again", preferredStyle: .alert)
                    
                    let okAction = UIAlertAction(title: "OK", style: .default, handler: nil)
                    alert.addAction(okAction)
                    
                    self?.present(alert, animated: true, completion: nil)
                    
                    
                    if let done = completion {
                        
                        done(error, true)
                    }
                } else if !(self?.exportSession?.cancelled ?? false) {
                    //success
                    if let done = completion {
                        done(nil, false)
                    }
                } else {
                    print("Export was cancelled")
                    if let done = completion {
                        done(nil, true)
                    }
                }
                    
                    
                })
            }
        
        
    }


  
}

// _Buttons
extension VideoPlayerViewController: SCAssetExportSessionDelegate {
    
    @objc func handleVideoSliderChange(_ slider: UISlider) {
        executeVideoSliderChange()
    }
    
    @objc func handleCutButton(_ button: UIButton) {
        print("cut")
    }
    
    @objc func handleRecordOverButton(_ button: UIButton) {
        print("record over")
    }
    
    @objc func handleMuteButton(_ button: UIButton) {
        
        if assetIsMuted {
            self.muteButton.isSelected = false
            playerView?.player?.volume = 1.0
            self.assetIsMuted = false
        } else {
            self.muteButton.isSelected = true
            playerView?.player?.volume = 0.0
            self.assetIsMuted = true
        }
    }
    
    @objc func handleStickersButton(_ button: UIButton) {
        showStickersViewController()
    }
    
    @objc func handleDrawButton(_ button: UIButton) {
        isDrawing = true
        canvasView.isUserInteractionEnabled = false
        doneButton.isHidden = false
        cancelButton.isHidden = true
        colorPickerView.isHidden = false
        undoButton.isHidden = false
        //hideToolbar(hide: true)
    }
    
    @objc func handleTextButton(_ button: UIButton) {
        executeTextButtonAction()
        
    }
    
    func executeTextButtonAction() {
        doneButton.isHidden = false
        cancelButton.isHidden = true
        colorPickerView.isHidden = false
        
        isTyping = true
        let textView = UIView(frame: CGRect(x: 0, y: canvasView.center.y, width: UIScreen.main.bounds.width, height: 50))
        
        textView.backgroundColor = UIColor.clear
        let subTextView = UITextView()
        
        subTextView.translatesAutoresizingMaskIntoConstraints = false
        subTextView.textAlignment = .center
        subTextView.font = UIFont(name: "Helvetica", size: 30)
        subTextView.textColor = textColor
        subTextView.layer.shadowColor = UIColor.black.cgColor
        subTextView.layer.shadowOffset = CGSize(width: 1.0, height: 0.0)
        subTextView.layer.shadowOpacity = 0.3
        subTextView.layer.shadowRadius = 1.0
        subTextView.layer.backgroundColor = backgroundTextColor.cgColor
        subTextView.autocorrectionType = .no
        subTextView.isScrollEnabled = false
        subTextView.delegate = self
        subTextView.tintColor = UIColor.white
        //subTextView.clipsToBounds = false
        
        self.canvasView.addSubview(textView)
        textView.addSubview(subTextView)
        
        
        NSLayoutConstraint.activate([
            subTextView.centerXAnchor.constraint(equalTo: textView.centerXAnchor),
            subTextView.centerYAnchor.constraint(equalTo: textView.centerYAnchor),
            ])
        
        
        addGestures(view: textView)
        subTextView.becomeFirstResponder()
        
    }
    
    @objc func handleDoneButton(_ button: UIButton) {
        executeDoneAction()
    }
    
    func executeDoneAction() {
        view.endEditing(true)
        doneButton.isHidden = true
        colorPickerView.isHidden = true
        cancelButton.isHidden = false
        undoButton.isHidden = true
        canvasView.isUserInteractionEnabled = true
        //hideToolbar(hide: false)
        isDrawing = false
    }


    
    @objc func handleUndoButton(_ button: UIButton) {
        //clear drawing
        drawings.popLast()
        if drawings.isEmpty {
            canvasView.image = nil
        } else {
            canvasView.image = drawings.last!
        }
        
        //clear stickers and textviews
//        for subview in canvasView.subviews {
//            subview.removeFromSuperview()
//        }
        
    }
    
    func countMetaElements(view: UIView) -> [String]? {
        var text: [String] = []
        for subview in view.subviews {
            if !subview.subviews.isEmpty {
                if let textSubview = subview as? UITextView {
                    text.append(textSubview.text)
                }
            }
        }
        if text.isEmpty {
            return nil
        } else {
            return text
        }
    }
    
    @objc func handleCompleteButton(_ button: UIButton) {
        //photoEditorDelegate?.doneEditing(image: img)
        
        
        let metaText = countMetaElements(view: canvasView)
        var drawingCount: Int?
        if !self.drawings.isEmpty {
            drawingCount = self.drawings.count
        }
        
        var volume = 1.0
        if self.muteButton.isSelected {
            volume = 0.0
        }
        
        
        
        
        if self.asset != nil && outputUrl != nil {
            
            
            self.completeButton.isEnabled = false
            
            let exportDispatch = DispatchGroup()
            exportDispatch.enter()
            
            let imgGenerator = AVAssetImageGenerator(asset: self.asset!)
            imgGenerator.appliesPreferredTrackTransform = true
            var thumbnail: UIImage?
            do {
                thumbnail = try canvasView.toImage().overlayWith(image: UIImage(cgImage: imgGenerator.copyCGImage(at: CMTime.zero, actualTime: nil)), posX: 0, posY: 0)
            } catch {
                print("Could not generate thumbnail")
            }
            
            var describingVideoOrientation: String?
            if let assetOrientation = self.asset?.videoOrientation() {
                switch assetOrientation {
                case .unknown:
                    describingVideoOrientation = "unknown"
                case .landscapeLeft:
                    describingVideoOrientation = "landscapeLeft"
                case .landscapeRight:
                    describingVideoOrientation = "landscapeRight"
                case .portrait:
                    describingVideoOrientation = "portrait"
                case .portraitUpsideDown:
                    describingVideoOrientation = "portraitUpsideDown"
                default:
                    break
                }
                
            }
            

//            self.navigationController?.pushViewController(FinishUploadPostViewController(videoUrl: self.outputUrl!, cleanedCompletionRatios: self.cleanedCompletionRatios, thumbnail: thumbnail, metaText: metaText, drawingCount: drawingCount, volume: volume, duration: self.asset?.duration.seconds, orientation: describingVideoOrientation, style: self.style!, exportDispatch: exportDispatch), animated: true)

            
            executeExport { [weak self] (error, cancelled) in
                if !cancelled {
                    self?.completeButton.isEnabled = true
                    exportDispatch.leave()
                } else if let error = error {
                    print("Could not execute Export with error: \(error)")
                }
            }
        } else if self.capturedImage != nil {
                self.completeButton.isEnabled = false
            
            let finalImage = self.captureImageView!.toImage()
            var describingImageOrientation: String?
            
            switch finalImage.imageOrientation {
            case .down:
                describingImageOrientation = "down"
            case .downMirrored:
                describingImageOrientation = "downMirrored"
            case .left:
                describingImageOrientation = "left"
            case .leftMirrored:
                describingImageOrientation = "leftMirrored"
            case .right:
                describingImageOrientation = "right"
            case .rightMirrored:
                describingImageOrientation = "rightMirrored"
            case .up:
                describingImageOrientation = "up"
            case .upMirrored:
                describingImageOrientation = "upMirrored"
            default:
                break
            }
//            self.navigationController?.pushViewController(FinishUploadPostViewController(finalImage: finalImage, orientation: describingImageOrientation, metaText: metaText, drawingCount: drawingCount, style: self.style!), animated: true)
             self.completeButton.isEnabled = true

            
        }
    }
    
    
    
    @objc internal func handleSaveButton(_ button: UIButton) {
        if self.asset != nil && !self.saveButton.isSelected {
            self.saveButton.isEnabled = false
            executeExport { [weak self] (error, cancelled) in
                 if !cancelled {
                    let saveToCameraRoll = SCSaveToCameraRollOperation()
                    saveToCameraRoll.saveVideoURL(self?.outputUrl, completion: { (message, error) in
                        
                        self?.saveButton.isEnabled = true
                        if let err = error {
                            print("Could not execute saveToCameraRoll with error: \(err)")
                            
                        } else {
                            self?.saveButton.isSelected = true
                            
                            print("successfully saved with message: \(message)")
                            
                        }
                    })

                 } else if let error = error {
                    print("Could not execute Export with error: \(error)")
                }
                
            }
        } else if self.capturedImage != nil {
            captureImageView?.toImage().saveToCameraRoll(completion: { (error) in
                if let error = error {
                    print("Could not save canvas to Camera Roll with error: \(error)")
                } else {
                    self.saveButton.isSelected = true
                    print("Successfully saved canvas to Camera Roll")
                }
            })
        }
    }
    
    
    
    @objc internal func handleCancelButton(_ button: UIButton) {
        self.playerView?.player?.pause()
        
        self.dismiss(animated: false, completion: nil)
    }
    
    //MAKR: helper methods
    
    @objc func image(_ image: UIImage, withPotentialError error: NSErrorPointer, contextInfo: UnsafeRawPointer) {
        let alert = UIAlertController(title: "Image Saved", message: "Image successfully saved to Photos library", preferredStyle: UIAlertController.Style.alert)
        alert.addAction(UIAlertAction(title: "OK", style: UIAlertAction.Style.default, handler: nil))
        self.present(alert, animated: true, completion: nil)
    }
    
    func hideControls() {

        textButton.isHidden = true
        drawButton.isHidden = true
        stickerButton.isHidden = true
        cutButton.isHidden = true
        recordOverButton.isHidden = true
        muteButton.isHidden = true
        saveButton.isHidden = true
        cutButton.isHidden = true
        cutButton.isHidden = true
    }
       
    
}



/**
 - didSelectView
 - didSelectImage
 */
protocol StickersViewControllerDelegate: AnyObject {
    /**
     - Parameter view: selected view from StickersViewController
     */
    func didSelectView(view: UIView)
    /**
     - Parameter image: selected Image from StickersViewController
     */
    func didSelectImage(image: UIImage)
    /**
     StickersViewController did Disappear
     */

}

/**
 - didSelectColor
 */
protocol ColorDelegate: AnyObject  {
    func changeColor(textColor: UIColor, backgroundColor: UIColor)
}

